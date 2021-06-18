package infrastucture

import (
	"context"
	"fmt"
	"time"

	"github.com/microservices/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const coefTimeout time.Duration = 10

func GetMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), coefTimeout*time.Second)
	defer cancel()
	config, err := utils.LoadConfig(".")
	if err != nil {
		return nil, err
	}
	URIRepo := config.DBDriver + "://" + config.DBHost + "/" + config.DBName
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URIRepo))

	if err != nil {
		return nil, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), coefTimeout*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	fmt.Println("Connect DB successfully")

	return client, nil
}
