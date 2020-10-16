package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func GetPostsCount(c *gin.Context) {
	var count int

	err := db.DB.QueryRow(`select count(id) from posts`).Scan(&count)

	if err != nil {
		utils.HandleError([]string{"some error"}, err.Error(), c, 500)
	}

	utils.SendPosRes(c, 200, gin.H{
		"count": count,
	})
}
