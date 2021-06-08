package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	app "microservice/application"
	"microservice/domain/service"
	infra "microservice/infrastructure"
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
		app.MakeCreateUserEndpoint(svc),
		app.DecodeCreateUserRequest,
		app.EncodeResponse,
	)
	GetUserByIdHandler := httptransport.NewServer(
		app.MakeGetUserByIdEndpoint(svc),
		app.DecodeGetUserByIdRequest,
		app.EncodeResponse,
	)

	DeleteUserHandler := httptransport.NewServer(
		app.MakeDeleteUserEndpoint(svc),
		app.DecodeDeleteUserRequest,
		app.EncodeResponse,
	)
	UpdateUserHandler := httptransport.NewServer(
		app.MakeUpdateUserEndpoint(svc),
		app.DecodeUpdateUserRequest,
		app.EncodeResponse,
	)

	r := mux.NewRouter()
	http.Handle("/", r)
	http.Handle("/create", CreateUserHandler)
	r.Handle("/users/{id}", GetUserByIdHandler)
	r.Handle("/deleteuser/{id}", DeleteUserHandler)
	r.Handle("/updateuser", UpdateUserHandler)
	http.ListenAndServe(":8080", nil)
}
