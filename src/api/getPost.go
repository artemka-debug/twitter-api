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

	var comments []map[string]interface{}

	errorGetting := db.DB.QueryRow(`select nickname, title, time, text, likes, user_id from posts where id = ?`, id).Scan(&nickname, &title, &time, &text, &likes, &userId)
	res, err := db.DB.Query(`select user_id, text, nickname from comments where post_id = ?`, id)

	if errorGetting != nil || err != nil {
		utils.HandleError([]string{"this post is no longer available"}, errorGetting.Error(), c, 400)
		return
	}

	for res.Next() {
		var (
			userId   int
			text     string
			nickname string
		)
		comment := make(map[string]interface{})

		if err := res.Scan(&userId, &text, &nickname); err != nil {
			utils.HandleError([]string{"could not get comments"}, err.Error(), c, 500)
			return
		}

		comment["userId"] = userId
		comment["text"] = text
		comment["nickname"] = nickname

		comments = append(comments, comment)
	}

	utils.SendPosRes(c, 200, gin.H{
		"nickname": nickname,
		"title":    title,
		"text":     text,
		"time":     time,
		"likes":    likes,
		"userId":   userId,
		"comments": comments,
	})
}
