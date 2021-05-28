package users

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/ibrahim-akrab/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/date_utils"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/errors"
)

const (
	indexUniqueEmail = "email_UNIQUE"
	noRowsResult     = "no rows in result set"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users where id = ?"
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
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), noRowsResult) {
			return errors.NewBadRequestError(err)
		}
		mysql.MySQLError
		return errors.NewInternalServerError(err)
	}
	return nil
}
