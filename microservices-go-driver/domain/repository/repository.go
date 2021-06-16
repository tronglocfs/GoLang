package repository

import (
	"context"

	"github.com/microservices/domain/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int) (model.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user *model.User) error
}
