package repository

import (
	"context"
	"microservice/domain/model"
)

type Repository interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	UpdateUser(ctx context.Context, id int, user model.User) error
}
