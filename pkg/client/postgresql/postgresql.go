package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go-rest-api/internal/config"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, cfg config.PostgresStorageConfig) (*pgxpool.Pool, error) {
	if maxAttempts <= 0 {
		return nil, fmt.Errorf("maxAttempts must be greater than 0")
	}

	originalAttempts := maxAttempts
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	for maxAttempts > 0 {
		newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		pool, err := pgxpool.New(newCtx, dsn)
		cancel()

		if err != nil {
			fmt.Println("connection failed, retrying...")
			maxAttempts--
			time.Sleep(5 * time.Second)
			continue
		}

		return pool, nil
	}

	return nil, fmt.Errorf("connection failed after %d attempts", originalAttempts)
}
