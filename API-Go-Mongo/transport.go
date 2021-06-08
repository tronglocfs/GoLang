package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"
)

type createUserRequest struct {
	user User
}

type createUserResponse struct {
	Msg string `json:"msg"`
	Err string `json:"err,omitempty"`
}

type getUserByIdRequest struct {
	Id string `json:"id"`
}

type getUserByIdResponse struct {
	Data interface{} `json:"data"`
	Err  string      `json:"err,omitempty"`
}

type deleteUserRequest struct {
	Id string `json:"id"`
}

type deleteUserResponse struct {
	Msg string `json:"msg"`
	Err string `json:"err,omitempty"`
}

type updateUserRequest struct {
	Id   string `json:"id"`
	user User
}

type updateUserResponse struct {
	Err string `json:"err,omitempty"`
}

func makeCreateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		msg, err := svc.CreateUser(ctx, req.user)

		if err != nil {
			return createUserResponse{"", err.Error()}, nil
		}

		return createUserResponse{msg, ""}, nil

	}
}

func makeGetUserByIdEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(getUserByIdRequest)
		id, er := strconv.Atoi(req.Id)

		if er != nil {
			return getUserByIdResponse{"", er.Error()}, nil
		}

		data, err := svc.GetUserById(ctx, id)

		if err != nil {
			return getUserByIdResponse{"", err.Error()}, nil
		}

		return getUserByIdResponse{data, ""}, nil

	}
}

func makeDeleteUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		id, er := strconv.Atoi(req.Id)

		if er != nil {
			return deleteUserResponse{"", er.Error()}, nil
		}

		msg, err := svc.DeleteUser(ctx, id)

		if err != nil {
			return deleteUserResponse{"", err.Error()}, nil
		}

		return deleteUserResponse{msg, ""}, nil

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

	vars := mux.Vars(r)

	request = getUserByIdRequest{
		Id: vars["id"],
	}
	return request, nil

}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteUserRequest

	vars := mux.Vars(r)

	request = deleteUserRequest{
		Id: vars["id"],
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
