package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	body := c.Keys["body"].(utils.LoginSchema)
	var password string
	var id int
	errorSelectingFromDb := db.DB.QueryRow(`select User_PK, password from users where email = ?`, body.Email).Scan(&id, &password)

	if errorSelectingFromDb != nil {
		utils.HandleError("you dont have an account, you need to sign up", c, 400)
		return
	}

	token := utils.CreateToken(body.Email, body.Password)
	if token == "" {
		utils.SendErrorRes(c, "could not create token", "", 500)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password)); err != nil {
		utils.HandleError("wrong password", c, 403)
		return
	}

	utils.SendPosRes(token, c, 200, id)
}
