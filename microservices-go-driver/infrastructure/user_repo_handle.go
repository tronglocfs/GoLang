package infrastucture

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/microservices/domain/model"
	"github.com/microservices/domain/repository"
	"github.com/microservices/utils"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

type repo struct {
	db *mongo.Client
}

func NewRepo(db *mongo.Client) (repository.Repository, error) {
	return &repo{
		db: db,
	}, nil
}

func (repo *repo) CreateUser(ctx context.Context, user *model.User) error {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return err
	}

	collection := repo.db.Database(config.DBName).Collection(config.DBCollection)

	_, cancel := context.WithTimeout(context.Background(), coefTimeout*time.Second)
	defer cancel()
	count, err := collection.CountDocuments(ctx, bson.D{primitive.E{Key: "$or", Value: bson.A{
		bson.D{primitive.E{Key: "email", Value: user.Email}},
		bson.D{primitive.E{Key: "_id", Value: user.Userid}}}}})

	if err != nil {
		fmt.Println("Creating user fails")
		return err
	}

	if count > 0 {
		fmt.Println("Email or UserID existed")
		return ErrAlreadyExists
	}

	_, err = collection.InsertOne(ctx, *user)

	if err != nil {
		fmt.Println("Creating user fails")
		return err
	}

	return nil
}

func (repo *repo) GetUserByID(ctx context.Context, id int) (model.User, error) {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return model.User{}, err
	}

	collection := repo.db.Database(config.DBName).Collection(config.DBCollection)

	var data model.User

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	_, cancel := context.WithTimeout(context.Background(), coefTimeout*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&data)
	if err == mongo.ErrNoDocuments {
		fmt.Println("User does not exist")
		return data, ErrNotFound
	}

	return data, nil
}

func (repo *repo) DeleteUser(ctx context.Context, id int) error {
	config, err := utils.LoadConfig(".")
	if err != nil {
		return err
	}

	collection := repo.db.Database(config.DBName).Collection(config.DBCollection)

	_, cancel := context.WithTimeout(context.Background(), coefTimeout*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, bson.D{primitive.E{Key: "_id", Value: id}})

	if err != nil {
		fmt.Println("Deleting user fails")
		return err
	}

	if res.DeletedCount == 0 {
		fmt.Println("Not found documents")
		return ErrNotFound
	}

	return nil
}

func (repo *repo) UpdateUser(ctx context.Context, user *model.User) error {
	var email model.User

	config, err := utils.LoadConfig(".")
	if err != nil {
		return err
	}

	collection := repo.db.Database(config.DBName).Collection(config.DBCollection)
	_, cancel := context.WithTimeout(context.Background(), coefTimeout*time.Second)
	defer cancel()

	// check exist of id, if id doesn't exist => create
	count, err := collection.CountDocuments(ctx, bson.D{primitive.E{Key: "_id", Value: user.Userid}})

	if err != nil {
		return err
	}
	if count == 0 {
		_, err = collection.InsertOne(ctx, *user)

		if err != nil {
			fmt.Println("Create user (because id does not exist) fails")
			return err
		}

		return nil
	}

	count, err = collection.CountDocuments(ctx, bson.D{primitive.E{Key: "email", Value: user.Email}})

	if err != nil {
		fmt.Println("Update user fails")
		return err
	}

	err = collection.FindOne(ctx, bson.D{primitive.E{Key: "email", Value: user.Email}}).Decode(&email)

	if err != nil {
		return err
	}
	if count > 1 || (count == 1 && user.Userid != email.Userid) {
		err = errors.New("email existed")
		fmt.Println("Email existed")
		return err
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": user.Userid},
		bson.M{"$set": bson.M{"email": user.Email, "password": user.Password, "phone": user.Phone}})

	if err != nil {
		fmt.Println("Updating user fails")
		return ErrNotFound
	}
	return nil
}
