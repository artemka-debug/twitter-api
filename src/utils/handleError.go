package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func HandleError(err string, c *gin.Context, code int) {
	fmt.Println("ERROR Error:", err)
	token := c.GetHeader("Authorization")
	if token != "" {
		token = strings.Split(c.GetHeader("Authorization"), " ")[1]
	}

	SendErrorRes(c, err, token, code)
}
