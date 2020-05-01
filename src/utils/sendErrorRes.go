package utils

import (
	"github.com/gin-gonic/gin"
)

func SendErrorRes(c *gin.Context, err, token string, code int) {
	c.JSON(code, gin.H{
		"token": token,
		"result": false,
		"error": err,
	})
}

func SendPosRes(token string, c *gin.Context, code int) {
	c.JSON(code, gin.H{
		"token": token,
		"result": true,
		"error": nil,
	})
}
