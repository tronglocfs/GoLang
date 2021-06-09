package infrastucture

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func GetMongoDB() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost/test"))

	/*defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()*/

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connect DB successfully")

	return client
}
