package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func RemovePost(c *gin.Context)  {
	id := c.Param("id")

	_, errorDeleting :=  db.DB.Exec(`delete posts from posts
												left join liked_posts l on posts.id = l.post_id
												left join comments c on posts.id = c.post_id
											where posts.id = ?`, id)

	if errorDeleting != nil {
		utils.HandleError([]string{"could not delete your post"}, errorDeleting.Error(), c, 500)
		return
	}

	utils.SendPosRes(c, 200, gin.H{})
}
