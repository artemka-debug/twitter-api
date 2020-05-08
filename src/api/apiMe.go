package api

import (
	"github.com/artemka-debug/twitter-api/src/secret"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"strings"
)

func Me(c *gin.Context) {
	if c.GetHeader("Authorization") == "" {
		utils.HandleError([]string{"token is not provided"}, "auth token is not provided, please provide token", c, 401)
		return
	}

	var pl utils.CustomPayload
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	_, err := jwt.Verify([]byte(token), secret.AppKey, &pl)

	if err != nil {
		utils.HandleError([]string{"you have trouble with trouble something(needs to be good error), try to re-login into your account"}, err.Error(), c, 401)
		return
	}

	utils.SendPosRes(c, 200, gin.H{
		"token": token,
	})
}
