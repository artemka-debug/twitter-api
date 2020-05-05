package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RemoveUser(c *gin.Context) {
	body := c.Keys["body"].(utils.RemoveUserSchema)
	var password string
	errorSelectingFromDb := db.DB.QueryRow(`select password from users where User_PK = ?`, body.Id).Scan(&password)

	if errorSelectingFromDb != nil {
		utils.HandleError("you dont have an account, you need to sign up", c, 400)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password)); err != nil {
		utils.HandleError("wrong password", c, 403)
		return
	}
	_, errorDeletingUser := db.DB.Query(`delete from users where User_PK = ?`, body.Id)

	if errorDeletingUser != nil {
		utils.HandleError("could not delete user, try again", c,  403)
	}

	utils.SendPosRes("", c, 200, -1)
}
