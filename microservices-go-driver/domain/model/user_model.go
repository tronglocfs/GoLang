package model

type User struct {
	Userid   int    `json:"userid" bson:"_id"`
	Email    string ` json:"email"`
	Password string ` json:"password" bson:"pass"`
	Phone    string ` json:"phone"`
}
