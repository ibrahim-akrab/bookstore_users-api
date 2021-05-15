package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement user creation")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement user retrieval")
}
