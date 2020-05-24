package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	token := c.Keys["token"]
	userId := c.Keys["userId"]
	userInfo := struct {
		Id       int    `json:"id"`
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
		Status   string `json:"status"`
	}{}

	errorGettingUserInfo := db.DB.QueryRow(`select id, nickname, email, status from users where id = ?`, userId).Scan(
		&userInfo.Id, &userInfo.Nickname, &userInfo.Email, &userInfo.Status)

	if errorGettingUserInfo != nil {
		utils.HandleError([]string{"no users found"}, errorGettingUserInfo.Error(), c, 400)
		return
	}

	utils.SendPosRes(c, 200, gin.H{
		"token": token,
		"user":  userInfo,
	})
}
