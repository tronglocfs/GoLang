package infrastucture

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"microservice/domain/model"
)

const UserCollection = "gotestuser"
const dbName = "User_Test"

type repo struct {
	db *mongo.Client
}

func NewRepo(db *mongo.Client) (model.Repository, error) {
	return &repo{
		db: db,
	}, nil
}

func (repo *repo) CreateUser(ctx context.Context, user model.User) (string, error) {
	collection := repo.db.Database(dbName).Collection(UserCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, user)
	//id := res.InsertedID

	if err != nil {
		fmt.Println("Creating user fails")
		return "", err
	} else {
		msg := "User Created: " + user.Email
		fmt.Println(msg)
		return msg, nil
	}
}

func (repo *repo) GetUserById(ctx context.Context, id int) (interface{}, error) {
	collection := repo.db.Database(dbName).Collection(UserCollection)

	var data model.User

	filter := bson.D{{"userid", id}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filter).Decode(&data)
	if err == mongo.ErrNoDocuments {
		fmt.Println("User does not exist")
		return data, err
	}

	return data, nil
}

func (repo *repo) DeleteUser(ctx context.Context, id int) (string, error) {
	collection := repo.db.Database(dbName).Collection(UserCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, bson.D{{"userid", id}})

	if err != nil {
		fmt.Println("Deleting user fails")
		return "", err
	} else if res.DeletedCount == 0 {
		fmt.Println("Not found documents")
		err = errors.New("no documents")
		return "", err
	} else {
		msg := "User deleted successfully"
		return msg, nil
	}
}

func (repo *repo) UpdateUser(ctx context.Context, id int, user model.User) error {
	collection := repo.db.Database(dbName).Collection(UserCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, bson.M{"userid": id}, bson.M{"$set": bson.M{"email": user.Email, "password": user.Password, "phone": user.Phone}})

	if err != nil {
		fmt.Println("Updating user fails")
		return err
	} else {
		return nil
	}

}