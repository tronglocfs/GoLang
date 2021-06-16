package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	transport "github.com/microservices/application/transport/http"
	infra "github.com/microservices/infrastructure"
)

func main() {

	//var svc service.UserService

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

	//svc = service.NewService()

	h := transport.MakeHTTPHandler(repository)

	log.Fatal(http.ListenAndServe(":8080", h))
}
