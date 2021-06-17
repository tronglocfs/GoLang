package endpoints

import (
	"context"

	"github.com/microservices/application/service"

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
	service service.UserService
}

// MakeUserEndpoints initializes all Go kit endpoints for the service
func MakeUserEndpoints(s service.UserService) UserEndpoints {
	return &userEndpoints{
		service: s,
	}
}

func (s userEndpoints) MakeCreateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.service.CreateUserService(ctx, request)
	}
}

func (s userEndpoints) MakeGetUserByIdEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.service.GetUserByIdService(ctx, request)
	}
}

func (s userEndpoints) MakeDeleteUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.service.DeleteUserService(ctx, request)

	}
}

func (s userEndpoints) MakeUpdateUserEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.service.UpdateUserService(ctx, request)
	}
}
