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
	body := c.Keys["body"].(*utils.ResetPasswordSchema)
	id := -1
	errorSelectingFromDb := db.DB.QueryRow(`select id from users where email = ?`, body.Email).Scan(&id)

	if errorSelectingFromDb != nil {
		utils.HandleError([]string{"you dont have an account, you need to sign up"}, errorSelectingFromDb.Error(), c, 400)
		return
	}
	newPass := utils.GeneratePassword()
	errorSendingEmail := smtp.SendMail(env.SmtpHost+":587",
		smtp.PlainAuth("", env.Email, env.Password, env.SmtpHost),
		env.Email, []string{body.Email}, []byte(fmt.Sprintf("Your new password is %s", newPass)))

	if errorSendingEmail != nil {
		utils.HandleError([]string{"could not send your new password to your email"}, errorSendingEmail.Error(), c, 500)
		return
	}
	fmt.Println(newPass)
	saltedBytes := []byte(newPass)
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)

	if errorHashing != nil {
		utils.HandleError([]string{"try again"}, errorHashing.Error(), c, 500)
		return
	}
	_, errorQuerying := db.DB.Query(`update users set password = ? where email = ?`, hashedBytes, body.Email)

	if errorQuerying != nil {
		utils.HandleError([]string{"could not connect to database, try again"}, errorQuerying.Error(), c, 500)
		return
	}
	utils.SendPosRes(c, 200, gin.H{})
}
