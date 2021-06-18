package service

import (
	"github.com/microservices/domain/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const minLengthPass int = 8
const maxLengthPass int = 16

func UserValidation(user *model.User) error {
	return validation.ValidateStruct(user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Userid, validation.Required),
		validation.Field(&user.Password, validation.Required, validation.Length(minLengthPass, maxLengthPass)))
}
