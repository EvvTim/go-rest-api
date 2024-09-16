package author

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"go-rest-api/internal/author"
	"go-rest-api/pkg/client/postgresql"
	"go-rest-api/pkg/logging"
	"go-rest-api/pkg/utils"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}

}

func (r repository) Create(ctx context.Context, author *author.Author) error {
	q := `
		INSERT INTO author 
		    (name)
		VALUES 
		    ($1) 
		RETURNING id
		`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
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

func (r repository) GetList(ctx context.Context) ([]*author.Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE is_deleted = false
		`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))
	rows, err := r.client.Query(ctx, q)

	if err != nil {
		return nil, err
	}

	authors := make([]*author.Author, 0)

	for rows.Next() {
		a := new(author.Author)

		if err := rows.Scan(&a.ID, &a.Name); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (r repository) GetByUUID(ctx context.Context, uuid string) (*author.Author, error) {
	q := `
		SELECT id, name * FROM public.author 
		                  WHERE id = $1
						  AND is_deleted = false
		`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	a := new(author.Author)

	if err := r.client.QueryRow(ctx, q, uuid).Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}

	return a, nil
}

func (r repository) Update(ctx context.Context, author *author.Author) error {
	q := `
		UPDATE public.author SET name = $2 WHERE id = $1
		`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	if _, err := r.client.Exec(ctx, q, author.ID, author.Name); err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, uuid string) error {
	q := ` UPDATE public.author SET is_deleted = true WHERE id = $1 `

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", utils.FormatQuery(q)))

	if _, err := r.client.Exec(ctx, q, uuid); err != nil {
		return err
	}

	return nil
}
