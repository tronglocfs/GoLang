package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/microservices/application/endpoints"
	httpapp "github.com/microservices/application/transport/http"
	"github.com/microservices/domain/service"
	infra "github.com/microservices/infrastructure"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	var svc service.UserService

	db, err := infra.GetMongoDB()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	repository, err := infra.NewRepo(db)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	svc = service.NewService(repository)

	CreateUserHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeCreateUserEndpoint(),
		httpapp.DecodeCreateUserRequest,
		httpapp.EncodeCreateUserResponse,
	)
	GetUserByIdHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeGetUserByIdEndpoint(),
		httpapp.DecodeGetUserByIdRequest,
		httpapp.EncodeGetUserByIDResponse,
	)

	DeleteUserHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeDeleteUserEndpoint(),
		httpapp.DecodeDeleteUserRequest,
		httpapp.EncodeDeleteUserResponse,
	)
	UpdateUserHandler := httptransport.NewServer(
		endpoints.MakeUserEndpoints(svc).MakeUpdateUserEndpoint(),
		httpapp.DecodeUpdateUserRequest,
		httpapp.EncodeUpdateUserResponse,
	)

	r := mux.NewRouter()
	http.Handle("/", r)
	r.Handle("/user-management/users", CreateUserHandler).Methods("POST") // POST
	r.Handle("/user-management/users/{id}", GetUserByIdHandler).Methods("GET")
	r.Handle("/user-management/users/{id}", DeleteUserHandler).Methods("DELETE")
	r.Handle("/user-management/users", UpdateUserHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
