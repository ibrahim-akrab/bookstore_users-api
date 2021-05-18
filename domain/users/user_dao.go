package users

import (
	"net/http"

	"github.com/ibrahim-akrab/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

// save the user to database
func (user *User) Save() *errors.RestErr {
	if usersDB[user.Id] != nil {
		return &errors.RestErr{Message: "user already exists", Status: http.StatusBadRequest}
	}
	usersDB[user.Id] = user

	return nil
}

// get a user from the database using primary key
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return &errors.RestErr{Message: "user doesn't exists", Status: http.StatusBadRequest}
	}
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
