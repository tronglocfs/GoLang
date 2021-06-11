package main

import (
	"net/http"
	"os"

	"microservice/application/endpoints"
	httpapp "microservice/application/transport/http"
	"microservice/domain/service"
	infra "microservice/infrastructure"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	var svc service.UserService

	db := infra.GetMongoDB()

	repository, err := infra.NewRepo(db)
	if err != nil {
		os.Exit(-1)
	}

	svc = service.NewService(repository)

	CreateUserHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeCreateUserEndpoint(),
		httpapp.DecodeCreateUserRequest,
		httpapp.EncodeResponse,
	)
	GetUserByIdHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeGetUserByIdEndpoint(),
		httpapp.DecodeGetUserByIdRequest,
		httpapp.EncodeResponse,
	)

	DeleteUserHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeDeleteUserEndpoint(),
		httpapp.DecodeDeleteUserRequest,
		httpapp.EncodeResponse,
	)
	UpdateUserHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeUpdateUserEndpoint(),
		httpapp.DecodeUpdateUserRequest,
		httpapp.EncodeResponse,
	)

	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/user-management/users", CreateUserHandler).Methods("POST") // POST
	r.Handle("/user-management/users/{id}", GetUserByIdHandler).Methods("GET")
	r.Handle("/user-management/users/{id}", DeleteUserHandler).Methods("DELETE")
	r.Handle("/user-management/users", UpdateUserHandler).Methods("PUT")
	http.ListenAndServe(":8080", nil)
}
