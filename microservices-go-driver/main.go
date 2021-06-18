package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/microservices/application/service"
	transport "github.com/microservices/application/transport/http"
	infra "github.com/microservices/infrastructure"
)

func main() {
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

	svc := service.NewUserService(repository)

	h := transport.MakeHTTPHandler(svc)

	log.Fatal(http.ListenAndServe(":8080", h))
}
