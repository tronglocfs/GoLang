package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
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
	GetByIdHandler := httptransport.NewServer(
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

	http.Handle("/create", CreateUserHandler)
	http.Handle("/user/1", GetByIdHandler)
	http.Handle("/deleteuser/1", DeleteUserHandler)
	http.Handle("/updateuser/", UpdateUserHandler)
	http.ListenAndServe(":8000", nil)
}
