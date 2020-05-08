package middleware

import (
	"errors"
	"fmt"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
	"strings"
)

func InputValidate(c *gin.Context) {
	var err error

	if c.Request.URL.Path == "/sign-up" {
		fields := c.Keys["body"].(utils.SignupSchema)

		err = validation.ValidateStruct(&fields,
				validation.Field(&fields.Email, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`))),
				validation.Field(&fields.Password, validation.Required, validation.Match(regexp.MustCompile(`\w{6,24}`))),
				validation.Field(&fields.Nickname, validation.Match(regexp.MustCompile(`[a-z1-9_'\-]{3,40}`))),
				validation.Field(&fields.Status, validation.Length(0, 150)))
	} else if c.Request.URL.Path == "/login" {
		fields := c.Keys["body"].(utils.LoginSchema)

		err = validation.ValidateStruct(&fields,
			validation.Field(&fields.Email, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`))),
			validation.Field(&fields.Password, validation.Required, validation.Match(regexp.MustCompile(`\w{6,24}`))))
	} else if c.Request.URL.Path == "/user/password/reset" {
		fields := c.Keys["body"].(utils.ResetPasswordSchema)

		err = validation.ValidateStruct(&fields,
			validation.Field(&fields.Email, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`))))
	} else if c.Request.URL.Path == "/tweet" {
		fields := c.Keys["body"].(utils.PostSchema)

		err = validation.ValidateStruct(&fields,
			validation.Field(&fields.Title, validation.Required, validation.Match(regexp.MustCompile(`[a-z_\-]{3,40}`))),
			validation.Field(&fields.Text, validation.Required, validation.Length(3, 255)))
	} else if c.Request.URL.Path == "/comment" {
		fields := c.Keys["body"].(utils.CommentSchema)

		err = validation.ValidateStruct(&fields,
			validation.Field(&fields.PostId, validation.Required),
			validation.Field(&fields.Text, validation.Required, validation.By(func(text interface{}) error {
				if strings.TrimSpace(text.(string)) == "" {
					return errors.New(" comment cannot be blank")
				}

				return nil
			}), validation.Length(1, 255)))
	} else if c.Request.URL.Path == "/user" {
		fields := c.Keys["body"].(utils.EditSchema)

		err = validation.ValidateStruct(&fields,
			validation.Field(&fields.Nickname, validation.Match(regexp.MustCompile(`[a-z1-9_'\-]{3,40}`))),
			validation.Field(&fields.Status, validation.Length(0, 150)))
	} else if c.Request.URL.Path == "/user/password" {
		fields := c.Keys["body"].(utils.ChangePassword)

		err = validation.ValidateStruct(&fields,
			validation.Field(&fields.Password, validation.Required, validation.Match(regexp.MustCompile(`\w{6,24}`))))
	}

	var errorsForUsers []string

	if err != nil {
		for _, v := range strings.Split(err.Error(), ";") {
			errorsForUsers = append(errorsForUsers, v)
		}

		fmt.Println(errorsForUsers)
		utils.HandleError(errorsForUsers, err.Error(), c, 400)
		c.Abort()
	}

	c.Next()
}
