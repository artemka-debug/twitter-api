package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/env"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"net/smtp"

	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
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

	url := "https://passwordwolf.com/api/?length=10&upper=off&lower=on&special=off&exclude=01234&repeat=1"
	method := "GET"

	client := &http.Client {}
	req, errorRequsting := http.NewRequest(method, url, nil)

	if utils.HandleError(errorRequsting, c) {
		return
	}

	res, errorDoing := client.Do(req)
	if utils.HandleError(errorDoing, c) {
		return
	}
	defer res.Body.Close()
	b, errorReading := ioutil.ReadAll(res.Body)

	if utils.HandleError(errorReading, c) {
		return
	}
	var data interface{}
	errorParsingJSON := json.Unmarshal(b, &data)
	if utils.HandleError(errorParsingJSON, c) {
		return
	}

	newPass := data.([]interface{})[0].(map[string]interface{})
	errorSendingEmail := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", env.Email, env.Password, "smtp.gmail.com"),
		env.Email, []string{body["email"].(string)}, []byte(fmt.Sprintf("Your new password is %s", newPass["password"].(string))))

	if utils.HandleError(errorSendingEmail, c) {
		return
	}
	fmt.Println(newPass)
	saltedBytes := []byte(newPass["password"].(string))
	hashedBytes, errorHashing := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)

	if utils.HandleError(errorHashing, c) {
		return
	}
	_, errorQuerying := db.DB.Query(`update users set password = ? where email = ?`, hashedBytes, body["email"])

	if utils.HandleError(errorQuerying, c) {
		return
	}
	utils.SendPosRes("", c)
}
