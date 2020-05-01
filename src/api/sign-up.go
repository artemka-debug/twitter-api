package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	body := c.Keys["body"].(utils.SignupSchema)
	saltedBytes := []byte(body.Password)
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if errorHashing != nil {
		utils.HandleError("try again", c, 500)
		return
	}

	// checking if user already is in db
	rows, errorQuerying := db.DB.Query("select email from Users where email = ?", body.Email)
	if errorQuerying != nil || rows.Err() != nil{
		utils.HandleError("could not connect to database, try again", c, 500)
		return
	}

	users := db.ReadSelect(rows)
	//// Pushing if user is not in db
	if len(users) != 0 {
		utils.HandleError("we already have an account with this email", c, 403)
		return
	}

	_, errorPushingUser := db.DB.Exec(`insert into Users (name, email, password, status, gender, totalLikes, profilePhoto)
													values (?, ?, ?, ?, '', 0, '')`, body.Nickname, body.Email, string(hashedBytes), body.Status)

	if errorPushingUser != nil {
		utils.HandleError("could not add you to database, try again", c, 500)
		return
	}
	token := utils.CreateToken(body.Email, body.Password)

	utils.SendPosRes(token, c, 201)
}
