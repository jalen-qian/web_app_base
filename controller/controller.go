package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ContextUserIdKey   = "userId"
	ContextUserNameKey = "userName"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
