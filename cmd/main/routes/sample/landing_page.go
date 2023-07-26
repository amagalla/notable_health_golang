package sample

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LandingPage (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}