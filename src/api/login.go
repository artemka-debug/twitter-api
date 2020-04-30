package api

import (
	"errors"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
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
	token := utils.CreateToken(body["email"].(string), body["password"].(string))
	if token == "" {
		utils.SendErrorRes(c, "Not Authorized", "")
	}

	if err := bcrypt.CompareHashAndPassword(user[0].([]byte), []byte(body["password"].(string))); err != nil {
		utils.HandleError(errors.New("wrong password"), c)
		return
	}

	utils.SendPosRes(token, c)
}
