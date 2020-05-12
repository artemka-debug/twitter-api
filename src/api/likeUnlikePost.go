package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func LikeUnlikePost(c *gin.Context) {
	postId := c.Param("id")
	userIdLike := c.Keys["userId"]
	var userIdPost int

	errorGetting := db.DB.QueryRow(`select user_id from Posts
                                      where id = ?`, postId).Scan(&userIdPost)

	if errorGetting != nil {
		utils.HandleError([]string{"this post not longer available"}, errorGetting.Error(), c, 400)
		return
	}

	alreadyLiked := []int{0, 0}
	like := 0
	var liked bool

	errorSelecting := db.DB.QueryRow(`select post_id, user_id from liked_posts where user_id = ? and post_id = ?`, userIdLike, postId).Scan(&alreadyLiked[0], &alreadyLiked[1])

	if errorSelecting != nil {
		// like tweet
		like = 1
		liked = true

		_, errorInserting := db.DB.Exec(`insert into liked_posts(user_id, post_id)
                                                values (?, ?)`, userIdLike, postId)

		if errorInserting != nil {
			utils.HandleError([]string{"could not like post"}, errorInserting.Error(), c, 500)
			return
		}
	} else {
		//unlike tweet
		like = -1
		liked = false

		_, errorDeleting := db.DB.Exec(`delete from liked_posts
                                                            where user_id = ? and post_id = ?`, userIdLike, postId)

		if errorDeleting != nil {
			utils.HandleError([]string{"could not dislike post"}, errorDeleting.Error(), c, 500)
			return
		}
	}

	_, errorLiking := db.DB.Exec(`update Posts
								set likes = likes + ?
								where id = ?`, like, postId)

	if errorLiking != nil {
		utils.HandleError([]string{"could not like post"}, errorLiking.Error(), c, 500)
	}

	utils.SendPosRes(c, 200, gin.H{
		"liked": liked,
	})
}