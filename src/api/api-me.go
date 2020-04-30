package api

import (
	"github.com/artemka-debug/twitter-api/src/secret"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
	"strings"
)

func Me(c *gin.Context) {
	var pl utils.CustomPayload
	token := strings.Split(c.GetHeader("Authorization"), " ")[1]
	_, err := jwt.Verify([]byte(token), secret.AppKey, &pl)
	if utils.HandleError(err, c) {
		return
	}

	utils.SendPosRes(token, c)
}