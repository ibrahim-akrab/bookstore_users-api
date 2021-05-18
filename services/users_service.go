package services

import (
	"github.com/ibrahim-akrab/bookstore_users-api/domain/users"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// TODO: implement user storage in database
	return &user, nil
}
