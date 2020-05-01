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
		utils.HandleError("token is not provided", c, 401)
		return
	}

	var pl utils.CustomPayload
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	_, err := jwt.Verify([]byte(token), secret.AppKey, &pl)
	if err != nil {
		utils.HandleError("authentication token is expired, try to re-login into your account", c, 401)
		return
	}

	utils.SendPosRes(token, c, 200)
}
