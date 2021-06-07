package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type createUserRequest struct {
	user User
}

type createUserResponse struct {
	Data string `json:"data"`
	Err  string `json:"err,omitempty"`
}

type getUserByIdRequest struct {
	Id int `json:"id"`
}

type getUserByIdResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err,omitempty"`
}

type deleteUserRequest struct {
	Id int `json:"id"`
}

type deleteUserResponse struct {
	Data string `json:"data"`
	Err  string `json:"err,omitempty"`
}

type updateUserRequest struct {
	Id   int `json:"id"`
	user User
}

type updateUserResponse struct {
	Err string `json:"err,omitempty"`
}

func makeCreateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		data, err := svc.CreateUser(ctx, req.user)

		if err != nil {
			return createUserResponse{"", err.Error()}, nil
		}

		return createUserResponse{data, ""}, nil

	}
}

func makeGetUserByIdEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByIdRequest)
		data, err := svc.GetUserById(ctx, req.Id)

		if err != nil {
			return getUserByIdResponse{"", err.Error()}, nil
		}

		return getUserByIdResponse{data, ""}, nil

	}
}

func makeDeleteUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		data, err := svc.DeleteUser(ctx, req.Id)

		if err != nil {
			return deleteUserResponse{"", err.Error()}, nil
		}

		return deleteUserResponse{data, ""}, nil

	}
}

func makeUpdateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUserRequest)
		err := svc.UpdateUser(ctx, req.user.Userid, req.user)

		if err != nil {
			return updateUserResponse{err.Error()}, nil
		}

		return updateUserResponse{""}, nil

	}
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request.user); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserByIdRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil

}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request updateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request.user); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
