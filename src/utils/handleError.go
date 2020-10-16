package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleError(errorForUser interface{}, err string, c *gin.Context, code int) {
	SendErrorRes(c, code, gin.H{
		"error":          err,
		"error_for_user": errorForUser,
	})
}
