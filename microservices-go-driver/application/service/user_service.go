package service

import (
	"context"
	"strconv"

	"github.com/microservices/domain/model"
	repo "github.com/microservices/domain/repository"
)

type UserService interface {
	CreateUserService(context.Context, interface{}) (interface{}, error)
	GetUserByIDService(context.Context, interface{}) (interface{}, error)
	DeleteUserService(context.Context, interface{}) (interface{}, error)
	UpdateUserService(context.Context, interface{}) (interface{}, error)
}

type userService struct {
	repoService repo.Repository
}

func NewUserService(r repo.Repository) UserService {
	return &userService{
		repoService: r,
	}
}

func (s userService) CreateUserService(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(CreateUserRequest)
	err := UserValidation(&req.User)

	if err != nil {
		return CreateUserResponse{err.Error()}, err
	}

	err = s.repoService.CreateUser(ctx, &req.User)

	if err != nil {
		return CreateUserResponse{err.Error()}, err
	}

	return CreateUserResponse{""}, nil
}

func (s userService) GetUserByIDService(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(GetUserByIDRequest)

	id, err := strconv.Atoi(req.ID)

	if err != nil {
		return GetUserByIDResponse{model.User{}, err.Error()}, err
	}

	data, err := s.repoService.GetUserByID(ctx, id)

	if err != nil {
		return GetUserByIDResponse{model.User{}, err.Error()}, err
	}

	return GetUserByIDResponse{data, ""}, nil
}

func (s userService) DeleteUserService(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(DeleteUserRequest)
	id, err := strconv.Atoi(req.ID)

	if err != nil {
		return DeleteUserResponse{err.Error()}, err
	}
	err = s.repoService.DeleteUser(ctx, id)
	if err != nil {
		return DeleteUserResponse{err.Error()}, err
	}
	return DeleteUserResponse{""}, nil
}

func (s userService) UpdateUserService(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(UpdateUserRequest)
	err := UserValidation(&req.User)

	if err != nil {
		return UpdateUserResponse{err.Error()}, err
	}

	err = s.repoService.UpdateUser(ctx, &req.User)
	if err != nil {
		return UpdateUserResponse{err.Error()}, err
	}
	return UpdateUserResponse{""}, nil
}

type CreateUserRequest struct {
	User model.User
}

type CreateUserResponse struct {
	Err string `json:"err,omitempty"`
}

type GetUserByIDRequest struct {
	ID string `json:"id"`
}

type GetUserByIDResponse struct {
	Data model.User `json:"data"`
	Err  string     `json:"err,omitempty"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	Err string `json:"err,omitempty"`
}

type UpdateUserRequest struct {
	User model.User
}

type UpdateUserResponse struct {
	Err string `json:"err,omitempty"`
}
