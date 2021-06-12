package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/microservices/application/endpoints"
)

func DecodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request.User); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.GetUserByIdRequest

	vars := mux.Vars(r)

	request = endpoints.GetUserByIdRequest{
		Id: vars["id"],
	}
	return request, nil

}

func DecodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.DeleteUserRequest

	vars := mux.Vars(r)

	request = endpoints.DeleteUserRequest{
		Id: vars["id"],
	}

	return request, nil
}

func DecodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoints.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request.User); err != nil {
		return nil, err
	}
	return request, nil
}
