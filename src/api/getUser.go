package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")

	var nickname, status, email string

	errorGetting := db.DB.QueryRow(`select nickname, status, email from users where id = ?`, id).Scan(&nickname, &status, &email)

	if errorGetting != nil {
		utils.HandleError([]string{"you don't have an account"}, errorGetting.Error(), c, 400)
		return
	}

	utils.SendPosRes(c, 200, gin.H{
		"nickname": nickname,
		"status":   status,
		"email":    email,
	})
}
