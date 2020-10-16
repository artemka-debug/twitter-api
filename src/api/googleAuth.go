package api

import (
	"encoding/json"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"strings"
)

func GoogleAuth(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)
	var body map[string]interface{}

	_ = json.Unmarshal(data, &body)
	defer c.Request.Body.Close()

	t, _ := utils.VerifyGoogleIdToken(body["id_token"].(string))

	if t.IssuedTo != "988530795253-t56qlidiqla5qnj2o6cvfbfeo3448ts8.apps.googleusercontent.com" {
		utils.HandleError([]string{"some error"}, "fixed", c, 400)
		return
	}

	var email string
	errorQuerying := db.DB.QueryRow("select email from Users where email = ?", email).Scan(&email)
	if errorQuerying != nil {
		var nickname, status string
		var id int

		_ = db.DB.QueryRow(`select nickname, status, id from users where email = ?`, email).Scan(&nickname, &status, &id)

		token := utils.CreateToken(id)

		utils.SendPosRes(c, 201, gin.H{
			"token":    token,
			"user_id":  id,
			"nickname": nickname,
			"status":   status,
		})
		return
	}

	email = t.Email
	password := t.UserId
	nickname := strings.Split(email, "@")[0]
	saltedBytes := []byte(password)
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if errorHashing != nil {
		utils.HandleError(utils.ErrorForUser{"password": "try again"}, errorHashing.Error(), c, 500)
		return
	}

	res, errorPushingUser := db.DB.Exec(`insert into Users (nickname, email, password, status)
													values (?, ?, ?, ?)`, nickname, email, string(hashedBytes), "")
	if errorPushingUser != nil {
		utils.HandleError(utils.ErrorForUser{"email": "could not add you to database, try again"}, errorPushingUser.Error(), c, 500)
		return
	}
	id, _ := res.LastInsertId()

	token := utils.CreateToken(int(id))

	utils.SendPosRes(c, 201, gin.H{
		"token":    token,
		"user_id":  int(id),
		"nickname": nickname,
		"status":   "",
	})
}
