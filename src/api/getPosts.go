package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []map[string]interface{}
	rows, errorGetting := db.DB.Query(`select nickname, title, time, text, likes, user_id from posts order by time`)

	if errorGetting != nil {
		utils.HandleError([]string{"you don't have an account"}, errorGetting.Error(), c, 400)
		return
	}

	for rows.Next() {
		var (
			nickname string
			title    string
			text     string
			time     string
			likes    int
			userId   int
		)
		post := make(map[string]interface{})

		if err := rows.Scan(&nickname, &title, &time, &text, &likes, &userId); err != nil {
			utils.HandleError([]string{"could not get all posts"}, err.Error(), c, 500)
			return
		}

		post["nickname"] = nickname
		post["title"] = title
		post["text"] = text
		post["time"] = time
		post["likes"] = likes
		post["userId"] = userId

		posts = append(posts, post)
	}

	utils.SendPosRes(c, 200, gin.H{
		"posts": posts,
	})
}
