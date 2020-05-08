package utils

import (
	"github.com/gin-gonic/gin"
)

func SendErrorRes(c *gin.Context, code int, data gin.H) {
	c.JSON(code, gin.H{
		"meta": gin.H{
			"status": code,
			"result": false,
		},
		"data": data,
	})
}

func SendPosRes(c *gin.Context, code int, data gin.H) {
	c.JSON(code, gin.H{
		"meta": gin.H{
			"status": code,
			"result": true,
		},
		"data": data,
	})
}
