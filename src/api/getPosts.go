package api

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"net/url"
	"sort"
	"strconv"
)

func GetPosts(c *gin.Context) {
	test, errorParsing := url.ParseQuery(c.Request.URL.String())

	if errorParsing != nil {
		utils.HandleError([]string{"could not get tweets"}, errorParsing.Error(), c, 400)
		return
	}

	if _, ok := test["climit"]; !ok {
		utils.HandleError([]string{"could not get comments"}, "limit on comments was not provided", c, 400)
		return
	}

	if _, ok := test["climit"]; !ok {
		utils.HandleError([]string{"could not get comments"}, "limit on posts was not provided", c, 400)
		return
	}

	tweetLimit, _ := strconv.Atoi(test["climit"][0])
	comLimit, _ := strconv.Atoi(test["climit"][0])

	var posts []map[string]interface{}
	rows, errorGetting := db.DB.Query(`select id, nickname, title, time, text, likes, user_id from posts order by time limit ?`, tweetLimit)
	defer rows.Close()

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
	}

	rows, errorComments := db.DB.Query(`select id, user_id, text, nickname, post_id from comments order by post_id limit ?`, comLimit)
	defer rows.Close()

	if errorComments != nil {
		utils.HandleError([]string{"could not get comments"}, errorComments.Error(), c, 500)
		return
	}

	var comments []map[string]interface{}
	for rows.Next() {
		var (
			userId   int
			text     string
			nickname string
			postId   int
			id       int
		)
		comment := make(map[string]interface{})

		if err := rows.Scan(&id, &userId, &text, &nickname, &postId); err != nil {
			utils.HandleError([]string{"could not get comments"}, err.Error(), c, 500)
			return
		}

		comment["userId"] = userId
		comment["text"] = text
		comment["nickname"] = nickname
		comment["postId"] = postId
		comment["id"] = id

		comments = append(comments, comment)
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

	fmt.Println(posts)

	utils.SendPosRes(c, 200, gin.H{
		"posts": posts,
	})
}
