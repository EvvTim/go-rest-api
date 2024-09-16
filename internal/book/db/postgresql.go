package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"go-rest-api/internal/book"
	"go-rest-api/pkg/client/postgresql"
	"go-rest-api/pkg/logging"
	"go-rest-api/pkg/utils"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r repository) Create(ctx context.Context, book *book.Book) error {
	q := `
		INSERT INTO book
		    (name)
		VALUES
			($1)
		RETURNING id
		`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	if err := r.client.QueryRow(ctx, q, book.Name).Scan(&book.ID); err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			newErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s",
				pgErr.Message,
				pgErr.Detail,
				pgErr.Where,
				pgErr.Code,
			)
			r.logger.Error(newErr)
			return nil
		}
		return err
	}
	return nil
}

func (r repository) GetByUUID(ctx context.Context, uuid string) (*book.Book, error) {
	p := `
		SELECT id, name FROM public.book WHERE id = $1 AND is_deleted = false
		`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(p)))

	a := new(book.Book)

	if err := r.client.QueryRow(ctx, p, uuid).Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}

	return a, nil
}

func (r repository) GetList(ctx context.Context) ([]*book.Book, error) {
	q := `
		SELECT id, name FROM public.book WHERE is_deleted = false
		`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	rows, err := r.client.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	books := make([]*book.Book, 0)

	for rows.Next() {
		a := new(book.Book)

		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, err
		}
		books = append(books, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r repository) Update(ctx context.Context, book *book.Book) error {
	q := `
		UPDATE public.book SET name = $2 WHERE id = $1 AND is_deleted = false
		`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	if _, err := r.client.Exec(ctx, q, book.ID, book.Name); err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, uuid string) error {
	q := `UPDATE public.book SET is_deleted = true WHERE id = $1 AND is_deleted = false `

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	if _, err := r.client.Exec(ctx, q, uuid); err != nil {
		return err
	}

	return nil
}
