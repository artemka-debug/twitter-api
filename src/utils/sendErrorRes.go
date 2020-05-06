package utils

import (
	"github.com/gin-gonic/gin"
)

func SendErrorRes(c *gin.Context, code int, data gin.H) {
	c.JSON(code, gin.H{
		"data": gin.H{
			"status": code,
			"result": false,
		},
		"meta": data,
	})
}

func SendPosRes(c *gin.Context, code int, data gin.H) {
	c.JSON(code, gin.H{
		"data": gin.H{
			"status": code,
			"result": true,
		},
		"meta": data,
	})
}
