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

		if c.Request.URL.Path == "/sign-up" {
			parseBody.BodySignup(data, c)
		} else if c.Request.URL.Path == "/login" {
			parseBody.BodyLogin(data, c)
		} else if c.Request.URL.Path == "/reset-password" {
			parseBody.BodyResetPassword(data, c)
		} else if c.Request.URL.Path == "/post" {
			parseBody.BodyPost(data, c)
		} else if c.Request.URL.Path == "/remove-user" {
			parseBody.BodyRemoveUser(data, c)
		} else if c.Request.URL.Path == "/add-comment" {
			parseBody.BodyComment(data, c)
		}

		c.Next()
}
