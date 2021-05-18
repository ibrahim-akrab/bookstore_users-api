package users

import (
	"fmt"
	"net/http"

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
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(result)

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement user retrieval")
}
