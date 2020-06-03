package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"sort"
)

func GetPostsByUser(c *gin.Context) {
	userId := c.Param("id")

	var posts []map[string]interface{}
	rows, errorGetting := db.DB.Query(`select id, nickname, title, time, text, likes, user_id from posts where user_id = ? order by time`, userId)

	if errorGetting != nil {
		utils.HandleError([]string{"you don't have an account"}, errorGetting.Error(), c, 400)
		return
	}

	var comments []map[string]interface{}

	for rows.Next() {
		var (
			nickname string
			title    string
			text     string
			time     string
			likes    int
			userId   int
			id       int
		)
		post := make(map[string]interface{})

		if err := rows.Scan(&id, &nickname, &title, &time, &text, &likes, &userId); err != nil {
			utils.HandleError([]string{"could not get all posts"}, err.Error(), c, 500)
			return
		}

		post["postId"] = id
		post["nickname"] = nickname
		post["title"] = title
		post["text"] = text
		post["time"] = time
		post["likes"] = likes
		post["userId"] = userId

		posts = append(posts, post)

		r, errorComments := db.DB.Query(`select id, user_id, text, nickname, post_id from comments where post_id = ?`, post["postId"])

		if errorComments != nil {
			utils.HandleError([]string{"could not get comments"}, errorComments.Error(), c, 500)
			return
		}

		for r.Next() {
			var (
				comUserId   int
				comText     string
				comNickname string
				postId      int
				comId       int
			)
			comment := make(map[string]interface{})

			if err := r.Scan(&comId, &comUserId, &comText, &comNickname, &postId); err != nil {
				utils.HandleError([]string{"could not get comments here"}, err.Error(), c, 500)
				return
			}

			comment["userId"] = comUserId
			comment["text"] = comText
			comment["nickname"] = comNickname
			comment["postId"] = postId
			comment["id"] = comId

			comments = append(comments, comment)
		}
	}


	sort.Slice(posts, func(i, j int) bool {
		return posts[i]["postId"].(int) < posts[j]["postId"].(int)
	})

	for i := 0; i < len(posts); i++ {
		c := make([]map[string]interface{}, 0)
		j := 0

		for ; j < len(comments) && posts[i]["postId"] == comments[j]["postId"]; j++ {
			c = append(c, comments[j])
		}
		comments = comments[j:]

		posts[i]["comments"] = c

	}

	if len(posts) == 0 {
		posts = make([]map[string]interface{}, 0)
	}

	utils.SendPosRes(c, 200, gin.H{
		"posts": posts,
	})
}
