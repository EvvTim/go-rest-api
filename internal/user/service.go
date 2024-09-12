package user

import (
	"context"
	"go-rest-api/pkg/logging"
)

type Service struct {
	store  Storage
	logger *logging.Logger
}

func (s Service) Create(ctx context.Context, user *User) (string, error) {
	//return s.store.Create(ctx, user)

	return "", nil
}
