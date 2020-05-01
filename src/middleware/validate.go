package middleware

import (
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func InputValidate(c *gin.Context) {
	if c.Request.URL.Path == "/sign-up" {
		fields := c.Keys["body"].(utils.SignupSchema)

		err := validation.ValidateStruct(&fields,
				validation.Field(&fields.Email, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`))),
				validation.Field(&fields.Password, validation.Required, validation.Match(regexp.MustCompile(`\w{6,24}`))),
				validation.Field(&fields.Nickname, validation.Required, validation.Match(regexp.MustCompile(`[a-z_\-]{3,40}`))),
				validation.Field(&fields.Status, validation.Required, validation.Length(3, 150)))

		if err != nil {
			utils.HandleError(err.Error(), c, 400)
			c.Abort()
		}
	} else if c.Request.URL.Path == "/login" {
		fields := c.Keys["body"].(utils.LoginSchema)

		err := validation.ValidateStruct(&fields,
			validation.Field(&fields.Email, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`))),
			validation.Field(&fields.Password, validation.Required, validation.Match(regexp.MustCompile(`\w{6,24}`))))

		if err != nil {
			utils.HandleError(err.Error(), c, 400)
			c.Abort()
		}
	} else if c.Request.URL.Path == "/reset-password" {
		fields := c.Keys["body"].(utils.ResetPassword)

		err := validation.ValidateStruct(&fields,
			validation.Field(&fields.Email, validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`))))

		if err != nil {
			utils.HandleError(err.Error(), c, 400)
			c.Abort()
		}
	}

	c.Next()
}
