package middleware

import (
	"github.com/artemka-debug/twitter-api/src/middleware/parseBody"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
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
