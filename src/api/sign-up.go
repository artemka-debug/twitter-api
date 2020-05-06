package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func SignUp(c *gin.Context) {
	body := c.Keys["body"].(utils.SignupSchema)
	saltedBytes := []byte(body.Password)
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if errorHashing != nil {
		utils.HandleError([]string{"try again"}, errorHashing.Error(), c, 500)
		return
	}

	var email string
	errorQuerying := db.DB.QueryRow("select email from Users where email = ?", body.Email).Scan(&email)
	if errorQuerying == nil {
		utils.HandleError([]string{"you already have an account"}, "user is in db", c, 400)
		return
	}

	var nickname string
	if body.Nickname == "" {
		nickname = strings.Split(body.Email, "@")[0]
	} else {
		nickname = body.Nickname
	}

	var id int
	_, errorPushingUser := db.DB.Query(`insert into Users (name, email, password, status, gender, totalLikes, profilePhoto)
													values (?, ?, ?, ?, '', 0, '')`, nickname, body.Email, string(hashedBytes), body.Status)

	_ = db.DB.QueryRow(`select User_PK from users where email = ?`, body.Email).Scan(&id)
	if errorPushingUser != nil {
		utils.HandleError([]string{"could not add you to database, try again"}, errorPushingUser.Error(), c, 500)
		return
	}
	token := utils.CreateToken(id, body.Password)

	utils.SendPosRes(c, 201, gin.H{
		"token": token,
		"user_id": id,
	})
}
