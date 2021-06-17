package http

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/microservices/application/endpoints"
	"github.com/microservices/application/service"
)

func MakeHTTPHandler(service service.UserService) http.Handler {
	r := mux.NewRouter()
	e := endpoints.MakeUserEndpoints(service)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/user-management/users").Handler(httptransport.NewServer(
		e.MakeCreateUserEndpoint(),
		DecodeCreateUserRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/user-management/users/{id}").Handler(httptransport.NewServer(
		e.MakeGetUserByIdEndpoint(),
		DecodeGetUserByIdRequest,
		encodeResponse,
		options...,
	))
	r.Methods("PUT").Path("/user-management/users").Handler(httptransport.NewServer(
		e.MakeUpdateUserEndpoint(),
		DecodeUpdateUserRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/user-management/users/{id}").Handler(httptransport.NewServer(
		e.MakeDeleteUserEndpoint(),
		DecodeDeleteUserRequest,
		encodeResponse,
		options...,
	))

	return r
}
