package service

import (
	"context"

	"github.com/microservices/domain/model"
	"github.com/microservices/domain/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) (string, error)
	GetUserById(ctx context.Context, id int) (model.User, error)
	DeleteUser(ctx context.Context, id int) (string, error)
	UpdateUser(ctx context.Context, id int, user *model.User) error
}

type userService struct {
	repository repository.Repository
}

func NewService(rep repository.Repository) UserService {
	return &userService{
		repository: rep,
	}
}

func (s userService) CreateUser(ctx context.Context, user *model.User) (string, error) {
	/*uuid, _ := uuid.NewV4()
	id := uuid.String()
	user.Id = id*/
	// userDetails := model.User{
	// 	Userid:   user.Userid,
	// 	Email:    user.Email,
	// 	Password: user.Password,
	// 	Phone:    user.Phone,
	// }

	msg, err := s.repository.CreateUser(ctx, user)

	if err != nil {
		return "", err
	}
	return msg, nil
}

func (s userService) GetUserById(ctx context.Context, id int) (model.User, error) {
	var data model.User
	data, err := s.repository.GetUserById(ctx, id)

	if err != nil {
		return data, err
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

func (s userService) UpdateUser(ctx context.Context, id int, user *model.User) error {
	err := s.repository.UpdateUser(ctx, id, user)

	if err != nil {
		return err
	}

	return nil
}
