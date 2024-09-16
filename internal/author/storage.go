package author

import "context"

type Repository interface {
	Create(ctx context.Context, author *Author) error
	GetByUUID(ctx context.Context, uuid string) (*Author, error)
	GetList(ctx context.Context) ([]*Author, error)
	Update(ctx context.Context, author *Author) error
	Delete(ctx context.Context, uuid string) error
}
