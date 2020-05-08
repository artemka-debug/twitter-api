package middleware

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/secret"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"strings"
)

func VerifyToken(c *gin.Context) {
	var pl utils.CustomPayload

	if c.GetHeader("Authorization") == "" {
		utils.HandleError([]string{"try to re-login"}, "authentication header was not provided", c, 401)
		c.Abort()
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	_, err := jwt.Verify([]byte(token), secret.AppKey, &pl)

	if err != nil {
		utils.HandleError([]string{"try to re-login"}, err.Error(), c, 401)
		c.Abort()
		return
	}

	fmt.Println("PAYLOAD ID", pl.Id)
	c.Set("userId", pl.Id)
	c.Next()
}
