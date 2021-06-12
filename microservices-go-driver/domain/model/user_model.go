package model

type User struct {
	Userid   int    `json:"userid"`
	Email    string ` json:"email"`
	Password string ` json:"password"`
	Phone    string ` json:"phone"`
}
