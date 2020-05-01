package api

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/env"
	"golang.org/x/crypto/bcrypt"
	"net/smtp"

	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func ResetPassword(c *gin.Context) {
	body := c.Keys["body"].(utils.ResetPassword)
	rows, errorSelectingFromDb := db.DB.Query(`select password from users where email = ?`, body.Email)
	if errorSelectingFromDb != nil {
		utils.HandleError("could not connect to database, try again", c, 500)
		return
	}

	user := db.ReadSelect(rows)
	if len(user) == 0 {
		utils.HandleError("you dont have an account, you need to sign up", c, 403)
		return
	}

	newPass := utils.GeneratePassword()
	fmt.Println(newPass)
	errorSendingEmail := smtp.SendMail(env.SmtpHost+ ":587",
		smtp.PlainAuth("", env.Email, env.Password, env.SmtpHost),
		env.Email, []string{body.Email}, []byte(fmt.Sprintf("Your new password is %s", newPass)))

	if errorSendingEmail != nil {
		utils.HandleError("could not send your new password to your email", c, 500)
		return
	}
	fmt.Println(newPass)
	saltedBytes := []byte(newPass)
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)

	if errorHashing != nil {
		utils.HandleError("try again", c, 500)
		return
	}
	_, errorQuerying := db.DB.Query(`update users set password = ? where email = ?`, hashedBytes, body.Email)

	if errorQuerying != nil {
		utils.HandleError("could not connect to database, try again", c, 500)
		return
	}
	utils.SendPosRes("", c, 201)
}
