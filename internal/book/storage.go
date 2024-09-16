package book

import "context"

type Repository interface {
	Create(ctx context.Context, book *Book) error
	GetByUUID(ctx context.Context, uuid string) (*Book, error)
	GetList(ctx context.Context) ([]*Book, error)
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, uuid string) error
}
