package services

import (
	"net/http"

	"github.com/ibrahim-akrab/bookstore_users-api/domain/users"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, &errors.RestErr{Message: "invalid user", Status: http.StatusBadRequest, Error: err}
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
