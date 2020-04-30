package api

import (
	"errors"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	body := c.Keys["body"].(map[string]interface{})
	saltedBytes := []byte(body["password"].(string))
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if utils.HandleError(errorHashing, c) {
		return
	}
	if utils.HandleError(utils.InputValidate(body), c) {
		return
	}

	// checking if user already is in db
	rows, errorQuerying := db.DB.Query("select email from Users where email = ?", body["email"])
	if utils.HandleError(errorQuerying, c) {
		return
	}
	if utils.HandleError(rows.Err(), c) {
		return
	}

	users := db.ReadSelect(rows)
	//// Pushing if user is not in db
	if len(users) != 0 {
		if utils.HandleError(errors.New("user is in db"), c) {
			return
		}
	}

	_, errorPushingUser := db.DB.Exec(`insert into Users (name, email, password, status, gender, totalLikes, profilePhoto)
													values (?, ?, ?, ?, '', 0, '')`, body["nickname"], body["email"], string(hashedBytes), body["status"])

	if utils.HandleError(errorPushingUser, c) {
		return
	}
	token := utils.CreateToken(body["email"].(string), body["password"].(string))

	utils.SendPosRes(token, c)
}
