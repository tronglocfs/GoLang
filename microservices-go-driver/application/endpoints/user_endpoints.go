package endpoints

import (
	"context"
	"strconv"

	"github.com/microservices/domain/model"
	repo "github.com/microservices/domain/repository"

	"github.com/go-kit/kit/endpoint"
)

// UserEndpoints holds all Go kit endpoints for the Repo
type UserEndpoints interface {
	MakeCreateUserEndpoint() endpoint.Endpoint
	MakeGetUserByIdEndpoint() endpoint.Endpoint
	MakeDeleteUserEndpoint() endpoint.Endpoint
	MakeUpdateUserEndpoint() endpoint.Endpoint
}

// UserEndpoints Struct to instance endpoints
type userEndpoints struct {
	repoDomainService repo.Repository
}

// MakeUserEndpoints initializes all Go kit endpoints for the Repo
func MakeUserEndpoints(repo repo.Repository) UserEndpoints {
	return &userEndpoints{
		repoDomainService: repo,
	}
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

func (s userEndpoints) MakeCreateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		err := UserValidation(&req.User)
		if err != nil {
			return CreateUserResponse{err.Error()}, err
		}
		err = s.repoDomainService.CreateUser(ctx, &req.User)

		if err != nil {
			return CreateUserResponse{err.Error()}, err
		}

		return CreateUserResponse{""}, nil

	}
}

func (s userEndpoints) MakeGetUserByIdEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetUserByIdRequest)
		id, er := strconv.Atoi(req.Id)

		if er != nil {
			return GetUserByIdResponse{model.User{}, er.Error()}, er
		}

		data, err := s.repoDomainService.GetUserById(ctx, id)
		//fmt.Println(data)
		if err != nil {
			return data, err
		}

		return GetUserByIdResponse{data, ""}, nil

	}
}

func (s userEndpoints) MakeDeleteUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		id, err := strconv.Atoi(req.Id)

		if err != nil {
			return DeleteUserResponse{err.Error()}, err
		}

		err = s.repoDomainService.DeleteUser(ctx, id)

		if err != nil {
			return DeleteUserResponse{err.Error()}, err
		}

		return DeleteUserResponse{""}, nil

	}
}

func (s userEndpoints) MakeUpdateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		err := UserValidation(&req.User)
		if err != nil {
			return UpdateUserResponse{err.Error()}, err
		}

		err = s.repoDomainService.UpdateUser(ctx, &req.User)
		if err != nil {
			return UpdateUserResponse{err.Error()}, err
		}

		return UpdateUserResponse{""}, nil

	}
}
