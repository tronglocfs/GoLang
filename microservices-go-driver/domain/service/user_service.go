package service

import (
	"context"
	"microservice/domain/model"
	"microservice/domain/repository"

	"github.com/gofrs/uuid"
)

type UserService interface {
	CreateUser(ctx context.Context, user model.User) (string, error)
	GetUserById(ctx context.Context, id int) (interface{}, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	UpdateUser(ctx context.Context, id int, user model.User) error
}

type userService struct {
	repository repository.Repository
}

func NewService(rep repository.Repository) UserService {
	return &userService{
		repository: rep,
	}
}

func (s userService) CreateUser(ctx context.Context, user model.User) (string, error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user.Id = id
	userDetails := model.User{
		Id:       user.Id,
		Userid:   user.Userid,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}

	msg, err := s.repository.CreateUser(ctx, userDetails)

	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s userService) GetUserById(ctx context.Context, id int) (interface{}, error) {
	var data interface{}
	data, err := s.repository.GetUserById(ctx, id)

	if err != nil {
		return "", err
	}

	return data, nil
}

func (s userService) DeleteUser(ctx context.Context, id int) (string, error) {

	data, err := s.repository.DeleteUser(ctx, id)

	if err != nil {
		return "", err
	}

	return data, nil
}

func (s userService) UpdateUser(ctx context.Context, id int, user model.User) error {
	err := s.repository.UpdateUser(ctx, id, user)

	if err != nil {
		return err
	}

	return nil
}
