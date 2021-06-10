package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"microservice/domain/model"
	"microservice/domain/service"
	"strconv"
)

// UserEndpoints holds all Go kit endpoints for the User service.
type UserEndpoints interface {
	MakeCreateUserEndpoint() endpoint.Endpoint
	MakeGetUserByIdEndpoint() endpoint.Endpoint
	MakeDeleteUserEndpoint() endpoint.Endpoint
	MakeUpdateUserEndpoint() endpoint.Endpoint
	//UserValidation() (string, error)
}

// UserEndpoints Struct to instance endpoints
type userEndpoints struct {
	userDomainService service.UserService
}

// MakeUserEndpoints initializes all Go kit endpoints for the User service.
func MakeUserEndpoints(s service.UserService) UserEndpoints {
	return &userEndpoints{
		userDomainService: s,
	}
}

type CreateUserRequest struct {
	User model.User
}

type CreateUserResponse struct {
	Msg string `json:"msg"`
	Err string `json:"err,omitempty"`
}

type GetUserByIdRequest struct {
	Id string `json:"id"`
}

type GetUserByIdResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err,omitempty"`
}

type DeleteUserRequest struct {
	Id string `json:"id"`
}

type DeleteUserResponse struct {
	Msg string `json:"msg"`
	Err string `json:"err,omitempty"`
}

type UpdateUserRequest struct {
	Id   string `json:"id"`
	User model.User
}

type UpdateUserResponse struct {
	Err string `json:"err,omitempty"`
}

func (s userEndpoints) MakeCreateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		err := UserValidation(req.User)
		if err != nil {
			 return CreateUserResponse{"", err.Error()}, nil
		}
		msg, err := s.userDomainService.CreateUser(ctx, req.User)

		if err != nil {
			return CreateUserResponse{"", err.Error()}, nil
		}

		return CreateUserResponse{msg, ""}, nil

	}
}

func (s userEndpoints) MakeGetUserByIdEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetUserByIdRequest)
		id, er := strconv.Atoi(req.Id)

		if er != nil {
			return GetUserByIdResponse{"", er.Error()}, nil
		}

		data, err := s.userDomainService.GetUserById(ctx, id)

		if err != nil {
			return GetUserByIdResponse{"", err.Error()}, nil
		}

		return GetUserByIdResponse{data, ""}, nil

	}
}

func (s userEndpoints) MakeDeleteUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteUserRequest)
		id, er := strconv.Atoi(req.Id)

		if er != nil {
			return DeleteUserResponse{"", er.Error()}, nil
		}

		msg, err := s.userDomainService.DeleteUser(ctx, id)

		if err != nil {
			return DeleteUserResponse{"", err.Error()}, nil
		}

		return DeleteUserResponse{msg, ""}, nil

	}
}

func (s userEndpoints) MakeUpdateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateUserRequest)
		err := UserValidation(req.User)
		if err != nil {
			return UpdateUserResponse{err.Error()}, nil
		}
		err = s.userDomainService.UpdateUser(ctx, req.User.Userid, req.User)

		if err != nil {
			return UpdateUserResponse{err.Error()}, nil
		}

		return UpdateUserResponse{""}, nil

	}
}
