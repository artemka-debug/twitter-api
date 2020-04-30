package utils

import (
	"github.com/gin-gonic/gin"
)

func SendErrorRes(c *gin.Context, err, token string) {
	c.JSON(200, gin.H{
		"token": token,
		"result": false,
		"error": err,
	})
}

func SendPosRes(token string, c *gin.Context) {
	c.JSON(200, gin.H{
		"token": token,
		"result": true,
		"error": nil,
	})
}
