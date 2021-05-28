package users

import (
	"net/http"
	"strings"

	"github.com/ibrahim-akrab/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/date_utils"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User)
)

// save the user to database
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(err)
		}
		return errors.NewInternalServerError(err)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(err)
	}

	user.Id = userId
	return nil
}

// get a user from the database using primary key
func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
