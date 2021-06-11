package infrastucture

import (
	"context"
	"fmt"
	"time"

	"microservice/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoDB() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	config, err := utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	URIRepo := config.DBDriver + "://" + config.DBHost + "/" + config.DBName
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URIRepo))

	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connect DB successfully")

	return client
}
