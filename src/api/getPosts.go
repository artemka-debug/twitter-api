package api

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"

	"sort"
	"strconv"
	"time"
)

func GetPosts(c *gin.Context) {
	test := c.Request.URL.Query()

	var tweetLimit, comLimit, page int

	if v, ok := test["tlimit"]; !ok {
		tweetLimit = 100
	} else {
		tweetLimit, _ = strconv.Atoi(v[0])
	}

	if v, ok := test["page"]; !ok {
		page = 1
	} else {
		page, _ = strconv.Atoi(v[0])
	}

	if v, ok := test["climit"]; !ok {
		comLimit = 100
	} else {
		comLimit, _ = strconv.Atoi(v[0])
	}

	fmt.Println(page, test)
	var posts []map[string]interface{}
	rows, errorGetting := db.DB.Query(`select id, nickname, title, time, text, likes, user_id from posts order by time desc limit ? offset ?`, tweetLimit, page * 100)
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

	rows, errorComments := db.DB.Query(`select id, user_id, text, nickname, post_id, time from comments order by post_id limit ?`, comLimit)
	defer rows.Close()

	if errorComments != nil {
		utils.HandleError([]string{"could not get comments"}, errorComments.Error(), c, 500)
		return
	}

	var comments []map[string]interface{}
	for rows.Next() {
		var (
			userId    int
			text      string
			nickname  string
			postId    int
			id        int
			timestamp string
		)
		comment := make(map[string]interface{})

		if err := rows.Scan(&id, &userId, &text, &nickname, &postId, &timestamp); err != nil {
			utils.HandleError([]string{"could not get comments"}, err.Error(), c, 500)
			return
		}

		var layout = "2006-01-02 15:04:05.999"
		t, _ := time.Parse(layout, timestamp)

		comment["userId"] = userId
		comment["text"] = text
		comment["nickname"] = nickname
		comment["postId"] = postId
		comment["id"] = id
		comment["time"] = t

		comments = append(comments, comment)
	}

	for i := 0; i < len(posts); i++ {
		c := make([]map[string]interface{}, 0)

		c = utils.Filter(comments, func(id int) bool {
			if id == posts[i]["postId"] {
				return true
			}
			return false
		})

		sort.SliceStable(c, func(i, j int) bool {
			return c[i]["time"].(time.Time).Sub(c[j]["time"].(time.Time)) < 0
		})

		if len(c) == 0 {
			posts[i]["comments"] = []utils.ErrorForUser{}
		} else {
			posts[i]["comments"] = c
		}
	}

	if len(posts) == 0 {
		posts = make([]map[string]interface{}, 0)
	}

	utils.SendPosRes(c, 200, gin.H{
		"posts": posts,
	})
}
