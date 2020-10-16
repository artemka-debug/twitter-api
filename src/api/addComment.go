package api

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	body := c.Keys["body"].(*utils.CommentSchema)
	userId := c.Keys["userId"].(int)
	nickname := ""

	errorQuerying := db.DB.QueryRow(`select nickname from users where id = ?`, userId).Scan(&nickname)

	if errorQuerying != nil {
		utils.HandleError([]string{"you do not have an account"}, errorQuerying.Error(), c, 500)
		return
	}

	res, errorInserting := db.DB.Exec(`insert into comments(text, user_id, post_id, nickname)
											values (?, ?, ?, ?)`, body.Text, userId, body.PostId, nickname)

	if errorInserting != nil {
		utils.HandleError(utils.ErrorForUser{"comment": "could not add your comment, try again"}, errorInserting.Error(), c, 500)
		return
	}

	commentId, _ := res.LastInsertId()
	userIdToSendNotification := 0
	_ = db.DB.QueryRow(`select user_id from posts where id = ?`, body.PostId).Scan(&userIdToSendNotification)

	if userIdToSendNotification > 0 {
		utils.SendNotification(userIdToSendNotification, fmt.Sprintf("%s commented your post", nickname))
	}

	utils.SendPosRes(c, 201, gin.H{
		"id":       commentId,
		"text":     body.Text,
		"nickname": nickname,
		"userId":   userId,
		"postId":   body.PostId,
	})
}
