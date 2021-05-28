package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ibrahim-akrab/bookstore_users-api/domain/users"
	"github.com/ibrahim-akrab/bookstore_users-api/services"
	"github.com/ibrahim-akrab/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.RestErr{
			Message: "Error parsing json request",
			Status:  http.StatusBadRequest,
			Error:   err,
		}
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr.Error)
		return
	}
	fmt.Println(result)

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		c.JSON(http.StatusBadRequest, userErr)
		return
	}
	user := &users.User{Id: userId}
	err := user.Get()
	if err != nil {
		c.JSON(err.Status, err.Error.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
