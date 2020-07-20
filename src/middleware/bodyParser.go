package middleware

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/middleware/parseBody"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"io/ioutil"
)

func BodyParser(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	if len(data) == 0 {
		utils.HandleError([]string{"no data was send"}, "req body is empty", c, 400)
		c.Abort()
		return
	}

	switch {
	case c.Request.URL.Path == "/sign-up":
		parseBody.BodySignup(data, c)

		break
	case c.Request.URL.Path == "/login":
		parseBody.BodyLogin(data, c)

		break
	case c.Request.URL.Path == "/user/password/reset":
		parseBody.BodyResetPassword(data, c)

		break
	case c.Request.URL.Path == "/tweet":
		parseBody.BodyPost(data, c)

		break
	case c.Request.URL.Path == "/comment":
		parseBody.BodyComment(data, c)

		break
	case c.Request.URL.Path == "/user" && c.Request.Method == "PUT":
		parseBody.BodyEdit(data, c)

		break
	case c.Request.URL.Path == "/user/password":
		parseBody.BodyChangePassword(data, c)

		break
	case c.Request.URL.Path == "/notification/subscribe":
		parseBody.BodySubscription(data, c)

		break
	}

	c.Next()
}

//func parse(data []byte, body interface{}, c *gin.Context) {
//	valType := reflect.TypeOf(body)
//	test := reflect.New(valType)
//
//	var json = jsoniter.ConfigCompatibleWithStandardLibrary
//	err := json.Unmarshal(data, test)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	c.Set("body", test)
//}
