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

func parseUserId(c *gin.Context) (int64, error) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return 0, err
	}
	return userId, nil
}

func Create(c *gin.Context) {
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
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr.Error)
		return
	}
	fmt.Println(result)

	c.JSON(http.StatusCreated, result)
}

func Get(c *gin.Context) {
	userId, userErr := parseUserId(c)
	if userErr != nil {
		return
	}
	user, err := services.UsersService.GetUser(userId)
	if err != nil {
		c.JSON(err.Status, err.Error)
		return
	}
	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	userId, userErr := parseUserId(c)
	if userErr != nil {
		return
	}

	var user users.User
	user.Id = userId
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

	partialUpdate := c.Request.Method == http.MethodPut

	result, restErr := services.UsersService.UpdateUser(user, partialUpdate)

	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	userId, userErr := parseUserId(c)
	if userErr != nil {
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.String(http.StatusOK, "deleted user successfully")
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
