package api

import (
	"errors"
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/env"
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
			env.Email,
			env.Password,
			env.Host,
		)


	err := smtp.SendMail(
			env.Host+":25",
			auth,
			env.Email,
			[]string{body["email"].(string)},
			[]byte("hi"),
		)

	if utils.HandleError(err, c) {
		return
	}

	utils.SendPosRes("", c)
}
