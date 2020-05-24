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
		utils.HandleError(utils.ErrorForUser{"password": "try again"}, errorHashing.Error(), c, 500)
		return
	}

	var email string
	errorQuerying := db.DB.QueryRow("select email from Users where email = ?", body.Email).Scan(&email)
	if errorQuerying == nil {
		utils.HandleError(utils.ErrorForUser{"email": "you already have an account"}, "user is in db", c, 400)
		return
	}

	var nickname string
	if body.Nickname == "" {
		nickname = strings.Split(body.Email, "@")[0]
	} else {
		nickname = body.Nickname
	}

	res, errorPushingUser := db.DB.Exec(`insert into Users (nickname, email, password, status)
													values (?, ?, ?, ?)`, nickname, body.Email, string(hashedBytes), body.Status)
	if errorPushingUser != nil {
		utils.HandleError(utils.ErrorForUser{"email": "could not add you to database, try again"}, errorPushingUser.Error(), c, 500)
		return
	}
	id, _ := res.LastInsertId()

	token := utils.CreateToken(int(id))

	utils.SendPosRes(c, 201, gin.H{
		"token":   token,
		"user_id": int(id),
	})
}
