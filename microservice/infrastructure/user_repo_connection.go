package infrastucture

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

func GetMongoDB() *mgo.Database {

	host := "MONGO_HOST"
	dbName := "User_Test"
	fmt.Println("Info DB:", host, dbName)
	session, err := mgo.Dial("mongodb://localhost/test")
	if err != nil {
		os.Exit(2)
	}
	db = session.DB(dbName)

	return db
}
