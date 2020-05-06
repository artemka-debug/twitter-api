package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	body := c.Keys["body"].(utils.CommentSchema)
	nickname := ""

	errorQuerying := db.DB.QueryRow(`select name from users where User_PK = ?`, body.UserId).Scan(&nickname)

	if errorQuerying != nil {
		utils.HandleError([]string{"you do not have an account"}, errorQuerying.Error(), c, 500)
		return
	}

	_, errorInserting := db.DB.Exec(`insert into comments(text, user_id, post_id, author)
											values (?, ?, ?, ?)`, body.Text, body.UserId, body.PostId, nickname)

	if errorInserting != nil {
		utils.HandleError([]string{"could not add your comment, try again"}, errorInserting.Error(), c, 500)
		return
	}

	utils.SendPosRes(c, 201, gin.H{})
}