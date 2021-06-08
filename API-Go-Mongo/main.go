package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	var svc UserService

	db := GetMongoDB()

	repository, err := NewRepo(db)
	if err != nil {
		os.Exit(-1)
	}

	svc = NewService(repository)

	CreateUserHandler := httptransport.NewServer(
		makeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeResponse,
	)
	GetUserByIdHandler := httptransport.NewServer(
		makeGetUserByIdEndpoint(svc),
		decodeGetUserByIdRequest,
		encodeResponse,
	)

	DeleteUserHandler := httptransport.NewServer(
		makeDeleteUserEndpoint(svc),
		decodeDeleteUserRequest,
		encodeResponse,
	)
	UpdateUserHandler := httptransport.NewServer(
		makeUpdateUserEndpoint(svc),
		decodeUpdateUserRequest,
		encodeResponse,
	)

	r := mux.NewRouter()
	http.Handle("/", r)
	http.Handle("/create", CreateUserHandler)
	r.Handle("/users/{id}", GetUserByIdHandler)
	r.Handle("/deleteuser/{id}", DeleteUserHandler)
	r.Handle("/updateuser", UpdateUserHandler)
	http.ListenAndServe(":8080", nil)
}
