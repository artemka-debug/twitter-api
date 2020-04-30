package utils

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"regexp"
)

func InputValidate(fields map[string]interface{}) error {
	for k, _ := range fields {
		if k == "email" {
			errorEmail := validation.Validate(fields["email"], validation.Required, validation.Match(regexp.MustCompile(`^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`)))
			if errorEmail != nil {
				return errorEmail
			}
		} else if k == "password" {
			errorPassword := validation.Validate(fields["password"], validation.Required, validation.Match(regexp.MustCompile(`\w{6,24}`)))
			if errorPassword != nil {
				return errorPassword
			}
		} else if k == "nickname" {
			errorNickname := validation.Validate(fields["nickname"], validation.Required, validation.Match(regexp.MustCompile(`[a-z_\-]{3,40}`)))
			if errorNickname != nil {
				return errorNickname
			}
		} else if k == "status" {
			errorsStatus := validation.Validate(fields["status"], validation.Length(3, 150))
			if errorsStatus != nil {
				return errorsStatus
			}

		}
	}

	return nil
}
