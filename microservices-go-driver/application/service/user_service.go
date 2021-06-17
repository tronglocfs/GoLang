package service

import (
	"context"
	"strconv"

	"github.com/microservices/domain/model"
	repo "github.com/microservices/domain/repository"
)

type UserService interface {
	CreateUserService(context.Context, interface{}) (interface{}, error)
	GetUserByIdService(context.Context, interface{}) (interface{}, error)
	DeleteUserService(context.Context, interface{}) (interface{}, error)
	UpdateUserService(context.Context, interface{}) (interface{}, error)
}

type userService struct {
	repoService repo.Repository
}

func NewUserService(repo repo.Repository) UserService {
	return &userService{
		repoService: repo,
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

func (s userService) GetUserByIdService(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(GetUserByIdRequest)

	id, err := strconv.Atoi(req.Id)

	if err != nil {
		return GetUserByIdResponse{model.User{}, err.Error()}, err
	}

	data, err := s.repoService.GetUserById(ctx, id)

	if err != nil {
		return GetUserByIdResponse{model.User{}, err.Error()}, err
	}

	return GetUserByIdResponse{data, ""}, nil
}

func (s userService) DeleteUserService(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(DeleteUserRequest)
	id, err := strconv.Atoi(req.Id)

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

type GetUserByIdRequest struct {
	Id string `json:"id"`
}

type GetUserByIdResponse struct {
	Data model.User `json:"data"`
	Err  string     `json:"err,omitempty"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
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
