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
		utils.HandleError("authentication was not provided, try re-login into your account", c, 401)
		c.Abort()
		return
	}

	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	_, err := jwt.Verify([]byte(token), secret.AppKey, &pl)

	if err != nil {
		utils.HandleError("authentication token is expired, try re-login into your account", c, 401)
		c.Abort()
		return
	}

	// This code is probably checking if the token you are using belongs to you, because in payload there is email
	// of the user who created this token and if emails does not equal needs to abort
	// but for this i need to pass email every time i send and get some data
	//if pl.Email == c.Keys["body"].Email {
	//	utils.HandleError("authentication token is not yours, try re-login into your account", c, 401)
	//	c.Abort()
	//	return
	//}

	fmt.Println(pl)
	c.Next()
}
