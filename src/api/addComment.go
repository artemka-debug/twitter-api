package api

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	body := c.Keys["body"].(utils.CommentSchema)
	userId := c.Keys["userId"].(int)
	nickname := ""

	errorQuerying := db.DB.QueryRow(`select nickname from users where id = ?`, userId).Scan(&nickname)

	if errorQuerying != nil {
		utils.HandleError([]string{"you do not have an account"}, errorQuerying.Error(), c, 500)
		return
	}

	fmt.Println(body.Text, userId, body.PostId, nickname)

	_, errorInserting := db.DB.Exec(`insert into comments(text, user_id, post_id, nickname)
											values (?, ?, ?, ?)`, body.Text, userId, body.PostId, nickname)

	if errorInserting != nil {
		utils.HandleError([]string{"could not add your comment, try again"}, errorInserting.Error(), c, 500)
		return
	}

	utils.SendPosRes(c, 201, gin.H{})
}
