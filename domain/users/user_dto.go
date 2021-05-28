package users

import "github.com/go-playground/validator/v10"

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() error {
	validate := validator.New()
	return validate.Struct(user)
}
