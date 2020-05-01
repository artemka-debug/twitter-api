package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	body := c.Keys["body"].(utils.LoginSchema)
	rows, errorSelectingFromDb := db.DB.Query(`select password from users where email = ?`, body.Email)

	if errorSelectingFromDb != nil {
		utils.HandleError("could not connect to database, try again", c, 500)
		return
	}
	user := db.ReadSelect(rows)

	if len(user) == 0 {
		utils.HandleError("you dont have an account, you need to sign up", c,  403)
		return
	}
	token := utils.CreateToken(body.Email, body.Password)
	if token == "" {
		utils.SendErrorRes(c, "could not create token", "", 500)
	}

	if err := bcrypt.CompareHashAndPassword(user[0].([]byte), []byte(body.Password)); err != nil {
		utils.HandleError("wrong password", c, 403)
		return
	}

	utils.SendPosRes(token, c, 200)
}
