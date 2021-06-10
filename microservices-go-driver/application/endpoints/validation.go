package endpoints

import (
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"microservice/domain/model"
)

func UserValidation(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Userid, validation.Required),
		validation.Field(&user.Password, validation.Required, validation.Length(8, 16)))
}