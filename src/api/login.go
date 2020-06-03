package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	body := c.Keys["body"].(utils.LoginSchema)

	var password, nickname, status string
	var id int

	errorSelectingFromDb := db.DB.QueryRow(`select id, password, nickname, status from users where email = ?`, body.Email).Scan(&id, &password, &nickname, &status)

	if errorSelectingFromDb != nil {
		utils.HandleError(utils.ErrorForUser{"email": "you dont have an account, you need to sign up"}, errorSelectingFromDb.Error(), c, 400)
		return
	}

	token := utils.CreateToken(id)
	if token == "" {
		utils.SendErrorRes(c, 500, gin.H{
			"error": "could not create token",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password)); err != nil {
		utils.HandleError(utils.ErrorForUser{"password": "wrong password"}, err.Error(), c, 403)
		return
	}

	utils.SendPosRes(c, 200, gin.H{
		"token":   token,
		"user_id": id,
		"nickname": nickname,
		"status": status,
	})
}
