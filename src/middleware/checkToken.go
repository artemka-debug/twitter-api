package middleware

import (
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func CheckToken1(c *gin.Context) {
	if c.GetHeader("Token") != "hi" {
		c.JSON(200, utils.Res{
			Token: "hi",
			Result: false,
			Error: "nil",
		})
		return
	}

	c.JSON(200, utils.Res{
		Token: "hi",
		Result: true,
		Error: "",
	})
	c.Next()
}