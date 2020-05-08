package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c *gin.Context) {
	body := c.Keys["body"].(utils.ChangePassword)
	userId := c.Keys["userId"].(int)

	saltedBytes := []byte(body.Password)
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if errorHashing != nil {
		utils.HandleError([]string{"try again"}, errorHashing.Error(), c, 500)
		return
	}

	_, errorChanging := db.DB.Exec(`update Users set password = ? where id = ?`, string(hashedBytes), userId)

	if errorChanging != nil {
		utils.HandleError([]string{"could not change your password"}, errorChanging.Error(), c, 500)
		return
	}

	utils.SendPosRes(c, 200, gin.H{})
}
