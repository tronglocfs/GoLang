package model

type User struct {
	Id       string `json:"id"`
	Userid   int    `json:"userid"`
	Email    string ` json:"email"`
	Password string ` json:"password"`
	Phone    string ` json:"phone"`
}
