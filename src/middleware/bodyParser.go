package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"reflect"
)

func BodyParser(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	if len(data) == 0 {
		utils.HandleError([]string{"no data was send"}, "req body is empty", c, 400)
		c.Abort()
		return
	}

	var bodyType interface{}

	switch {
	case c.Request.URL.Path == "/sign-up":
		bodyType = new(utils.SignupSchema)

		break
	case c.Request.URL.Path == "/login":
		bodyType = new(utils.LoginSchema)

		break
	case c.Request.URL.Path == "/user/password/reset":
		bodyType = new(utils.ResetPasswordSchema)

		break
	case c.Request.URL.Path == "/tweet":
		bodyType = new(utils.PostSchema)

		break
	case c.Request.URL.Path == "/comment":
		bodyType = new(utils.CommentSchema)

		break
	case c.Request.URL.Path == "/user" && c.Request.Method == "PUT":
		bodyType = new(utils.EditSchema)

		break
	case c.Request.URL.Path == "/user/password":
		bodyType = new(utils.ChangePassword)

		break
	case c.Request.URL.Path == "/notification/subscribe":
		bodyType = new(utils.Subscription)

		break
	}

	t := reflect.TypeOf(bodyType)
	v := reflect.New(t.Elem())
	newP := v.Interface()
	err := json.Unmarshal(data, newP)

	if err != nil {
		fmt.Println(err)
	}

	c.Set("body", newP)
	c.Next()
}
