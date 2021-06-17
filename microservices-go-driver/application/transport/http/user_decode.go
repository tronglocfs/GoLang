package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/microservices/application/service"
)

func DecodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request.User); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.GetUserByIdRequest

	vars := mux.Vars(r)

	request = service.GetUserByIdRequest{
		Id: vars["id"],
	}
	return request, nil

}

func DecodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.DeleteUserRequest

	vars := mux.Vars(r)

	request = service.DeleteUserRequest{
		Id: vars["id"],
	}

	return request, nil
}

func DecodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request service.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request.User); err != nil {
		return nil, err
	}
	return request, nil
}
