package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
	id := c.Param("id")

	var nickname, title, text string
	var time string
	var likes, userId int

	errorGetting := db.DB.QueryRow(`select nickname, title, time, text, likes, user_id from posts where id = ?`, id).Scan(&nickname, &title, &time, &text, &likes, &userId)

	if errorGetting != nil {
		utils.HandleError([]string{"you don't have an account"}, errorGetting.Error(), c, 400)
		return
	}

	utils.SendPosRes(c, 200, gin.H{
		"nickname": nickname,
		"title": title,
		"text": text,
		"time": time,
		"likes": likes,
		"userId": userId,
	})
}
