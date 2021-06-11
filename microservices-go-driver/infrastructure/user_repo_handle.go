package infrastucture

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"microservice/config"
	"microservice/domain/model"
	"microservice/domain/repository"
)

type repo struct {
	db *mongo.Client
}

func NewRepo(db *mongo.Client) (repository.Repository, error) {
	return &repo{
		db: db,
	}, nil
}

func (repo *repo) CreateUser(ctx context.Context, user model.User) (string, error) {

	collection := repo.db.Database(config.DbName).Collection(config.UserCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	count, err := collection.CountDocuments(ctx, bson.D{{"$or", bson.A{
		bson.D{{"email", user.Email}},
		bson.D{{"userid", user.Userid}}}}})

	if err != nil {
		fmt.Println("Creating user fails")
		return "", err
	}

	if count > 0 {
		msg := "Email or UserID existed"
		fmt.Println("Email or UserID existed")
		return msg, err
	}

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		fmt.Println("Creating user fails")
		return "", err
	} else {
		msg := "User Created: " + user.Email
		fmt.Println(msg)
		return msg, nil
	}
}

func (repo *repo) GetUserById(ctx context.Context, id int) (model.User, error) {
	collection := repo.db.Database(config.DbName).Collection(config.UserCollection)

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
	collection := repo.db.Database(config.DbName).Collection(config.UserCollection)

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
	var email model.User
	collection := repo.db.Database(config.DbName).Collection(config.UserCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.D{{"email", user.Email}})

	if err != nil {
		fmt.Println("Update user fails")
		return err
	}

	err = collection.FindOne(ctx, bson.D{{"email", user.Email}}).Decode(&email)

	if count > 1 || (count == 1 && email.Email != user.Email) {
		err := errors.New("email existed")
		fmt.Println("Email existed")
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"userid": id}, bson.M{"$set": bson.M{"email": user.Email, "password": user.Password, "phone": user.Phone}})

	if err != nil {
		fmt.Println("Updating user fails")
		return err
	} else {
		return nil
	}

}
