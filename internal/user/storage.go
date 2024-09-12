package user

import "context"

type Storage interface {
	Create(ctx context.Context, user *User) (string, error)
	GetByUUID(ctx context.Context, uuid string) (*User, error)
	GetList(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, uuid string) error
}
