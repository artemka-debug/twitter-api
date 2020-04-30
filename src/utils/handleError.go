package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func HandleError(err interface{}, c *gin.Context) bool {
	switch err.(type) {
		case error:
			fmt.Println("ERROR Error:", err)
			token := c.GetHeader("Authorization")
			if token != "" {
				token = strings.Split(c.GetHeader("Authorization"), " ")[1]
			}

			SendErrorRes(c, err.(error).Error(), token)
			return true
	}

	return false
}
