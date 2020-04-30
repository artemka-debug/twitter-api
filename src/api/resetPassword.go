package api

import (
	"errors"
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"net/smtp"
)

func ResetPassword(c *gin.Context) {
	body := c.Keys["body"].(map[string]interface{})
	rows, errorSelectingFromDb := db.DB.Query(`select password from users where email = ?`, body["email"])
	if utils.HandleError(errorSelectingFromDb, c) {
		return
	}

	user := db.ReadSelect(rows)
	if len(user) == 0 {
		utils.HandleError(errors.New("no users found"), c)
		return
	}

	auth := smtp.PlainAuth(
			"",
			"dovgopoly123@gmail.com",
			"62137990Aa",
			"mail.example.com",
		)


	err := smtp.SendMail(
			"mail.example.com"+":25",
			auth,
			"dovgopoly123@gmail.com",
			[]string{body["email"].(string)},
			[]byte("hi"),
		)
	fmt.Println(auth, err)
}
