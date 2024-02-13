package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("Email is required"),
			validation.RuneLength(1, 30).Error("Email length must be between 1 and 30"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("Password is required"),
			validation.RuneLength(1, 30).Error("Password length must be between 1 and 30"),
		),
	)
}
