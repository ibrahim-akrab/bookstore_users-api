package services

import (
	"net/http"

	"github.com/ibrahim-akrab/bookstore_users-api/domain/users"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/crypto_utils"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, &errors.RestErr{Message: "invalid user", Status: http.StatusBadRequest, Error: err}
	}
	user.Password = crypto_utils.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	err := result.Get()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user users.User, partialUpdate bool) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)

	if err != nil {
		return nil, err
	}

	if partialUpdate {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
		if user.Password != "" {
			current.Password = crypto_utils.GetMD5(user.Password)
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Password = user.Password
		current.Status = user.Status
	}

	err = current.Update()
	if err != nil {
		return nil, err
	}
	return current, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func Search(status string) (*[]users.User, *errors.RestErr) {
	dao := &users.User{}
	users, err := dao.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return users, nil
}
