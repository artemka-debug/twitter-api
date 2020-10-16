package middleware

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"strings"
)

func InputValidate(c *gin.Context) {
	var err error

	switch {
	case c.Request.URL.Path == "/sign-up":
		fields := c.Keys["body"].(*utils.SignupSchema)

		err = validation.ValidateStruct(fields,
			validation.Field(&fields.Email, validation.Required, validation.By(utils.ValidateEmail)),
			validation.Field(&fields.Password, validation.Required, validation.By(utils.ValidatePassword)),
			validation.Field(&fields.Nickname, validation.Length(3, 40), validation.By(utils.ValidateNickname)),
			validation.Field(&fields.Status, validation.Length(0, 150)))

		break
	case c.Request.URL.Path == "/login":
		fields := c.Keys["body"].(*utils.LoginSchema)

		err = validation.ValidateStruct(fields,
			validation.Field(&fields.Email, validation.Required, validation.By(utils.ValidateEmail)),
			validation.Field(&fields.Password, validation.Required, validation.By(utils.ValidatePassword)))

		break
	case c.Request.URL.Path == "/user/password/reset":
		fields := c.Keys["body"].(*utils.ResetPasswordSchema)

		err = validation.ValidateStruct(fields,
			validation.Field(&fields.Email, validation.Required, validation.By(utils.ValidateEmail)))
	case c.Request.URL.Path == "/tweet":
		fields := c.Keys["body"].(*utils.TweetSchema)

		fmt.Println(fields)
		err = validation.ValidateStruct(fields,
			validation.Field(&fields.Title, validation.Required, validation.Length(1, 50), validation.By(utils.ValidateTitle)),
			validation.Field(&fields.Text, validation.Required, validation.Length(1, 255), validation.By(utils.ValidateTweetText)))

		break
	case c.Request.URL.Path == "/comment":
		fields := c.Keys["body"].(*utils.CommentSchema)

		err = validation.ValidateStruct(fields,
			validation.Field(&fields.PostId, validation.Required),
			validation.Field(&fields.Text, validation.Required, validation.Length(1, 255), validation.By(utils.ValidateCommentText)))

		break
	case c.Request.URL.Path == "/user" && c.Request.Method == "PUT":
		fields := c.Keys["body"].(*utils.EditSchema)

		err = validation.ValidateStruct(fields,
			validation.Field(&fields.Nickname, validation.By(utils.ValidateNickname)),
			validation.Field(&fields.Status, validation.Length(0, 150), validation.By(utils.ValidateStatus)))

		break
	case c.Request.URL.Path == "/user/password":
		fields := c.Keys["body"].(*utils.ChangePassword)

		err = validation.ValidateStruct(fields,
			validation.Field(&fields.Password, validation.Required, validation.By(utils.ValidatePassword)))

		break
	}

	errorsForUsers := make(utils.ErrorForUser, 0)

	if err != nil {
		for _, v := range strings.Split(err.Error(), ";") {
			var splitErr []string

			fmt.Println("ERROR", v)
			if v[0] == ' ' {
				splitErr = strings.Split(v[1:], ":")
			} else {
				splitErr = strings.Split(v, ":")
			}
			errorsForUsers[splitErr[0]] = splitErr[1][1:]
		}

		utils.HandleError(errorsForUsers, "Invalid input", c, 400)
		c.Abort()
	}

	c.Next()
}
