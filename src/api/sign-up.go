package api

import (
	"fmt"
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

	var email string
	// checking if user already is in db
	errorQuerying := db.DB.QueryRow("select email from Users where email = ?", body.Email).Scan(&email)
	if errorQuerying == nil {
		utils.HandleError("you already have an account", c, 500)
		return
	}

	var id int
	_, errorPushingUser := db.DB.Query(`insert into Users (name, email, password, status, gender, totalLikes, profilePhoto)
													values (?, ?, ?, ?, '', 0, '')`, body.Nickname, body.Email, string(hashedBytes), body.Status)

	_ = db.DB.QueryRow(`select User_PK from users where email = ?`, body.Email).Scan(&id)
	fmt.Println("ID", id)
	if errorPushingUser != nil {
		utils.HandleError("could not add you to database, try again", c, 500)
		return
	}
	token := utils.CreateToken(body.Email, body.Password)

	utils.SendPosRes(token, c, 201, id)
}
