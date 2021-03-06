package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func Post(c *gin.Context) {
	body := c.Keys["body"].(*utils.TweetSchema)
	userId := c.Keys["userId"].(int)
	var id int
	var nickname string
	errorSelectingFromDb := db.DB.QueryRow(`select id, nickname from users where id = ?`, userId).Scan(&id, &nickname)

	if errorSelectingFromDb != nil {
		utils.HandleError([]string{"you dont have an account, you need to sign up"}, errorSelectingFromDb.Error(), c, 400)
		return
	}
	res, errorPosting := db.DB.Exec(`insert into Posts(title, text, nickname, likes, user_id)
                                  values (?, ?, ?, ?, ?)`, body.Title, body.Text, nickname, 0, id)

	if errorPosting != nil {
		utils.HandleError([]string{"error posting, try again"}, errorPosting.Error(), c, 500)
		return
	}
	var timestamp string
	lastId, _ := res.LastInsertId()
	_ = db.DB.QueryRow(`select time from posts where id = ?`, lastId).Scan(&timestamp)

	utils.SendPosRes(c, 200, gin.H{
		"token":    strings.Split(c.GetHeader("Authorization"), " ")[1],
		"post_id":  int(lastId),
		"nickname": nickname,
		"title":    body.Title,
		"text":     body.Text,
		"time":     timestamp,
		"userId":   id,
		"likes":    0,
		"comments": []utils.ErrorForUser{},
	})
}
