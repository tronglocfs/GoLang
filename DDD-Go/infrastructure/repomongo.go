package infrastucture

import (
	"context"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"DDD-Go/domain/model"
	"DDD-Go/domain/service"
	"DDD-Go/application"
)
)

const UserCollection = "gotestuser"

type repo struct {
	db *mgo.Database
}

func NewRepo(db *mgo.Database) (Repository, error) {
	return &repo{
		db: db,
	}, nil
}

func (repo *repo) CreateUser(ctx context.Context, user User) (string, error) {
	err := db.C(UserCollection).Insert(user)
	if err != nil {
		fmt.Println("Error occured inside CreateUser in repo")
		return "", err
	} else {
		msg := "User Created: " + user.Email
		fmt.Println(msg)
		return msg, nil
	}
}

func (repo *repo) GetUserById(ctx context.Context, id int) (interface{}, error) {
	coll := db.C(UserCollection)
	data := []User{}
	err := coll.Find(bson.M{"userid": id}).Select(bson.M{}).All(&data)
	if err != nil {
		fmt.Println("Error occured inside GetUserById in repo")
		return "", err
	}
	return data, nil
}

func (repo *repo) DeleteUser(ctx context.Context, id int) (string, error) {
	coll := db.C(UserCollection)
	err := coll.Remove(bson.M{"userid": id})
	if err != nil {
		fmt.Println("Error occured inside delete in repo")
		return "", err
	} else {
		msg := "User deleted successfully"
		return msg, nil
	}
}

func (repo *repo) UpdateUser(ctx context.Context, id int, user User) error {
	coll := db.C(UserCollection)
	err := coll.Update(bson.M{"userid": id}, bson.M{"$set": bson.M{"email": user.Email, "password": user.Password, "phone": user.Phone}})
	if err != nil {
		fmt.Println("Error occured inside update user repo")
		return err
	} else {
		return nil
	}

}
