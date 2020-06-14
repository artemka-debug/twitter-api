package utils

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateStatus(textInterface interface{}) error {
	text := textInterface.(string)
	var err string

	if len(text) != 0 && strings.TrimSpace(text) == "" {
		err += "status cannot be made of only spaces"
	}

	if len(err) == 0 {
		return nil
	}

	return errors.New(err)
}

func ValidateTweetText(textInterface interface{}) error {
	text := textInterface.(string)
	var err string

	if strings.TrimSpace(text) == "" {
		err += "tweet text cannot be blank"
	}

	if len(err) == 0 {
		return nil
	}

	return errors.New(err)
}

func ValidateCommentText(textInterface interface{}) error {
	text := textInterface.(string)
	var err string

	if strings.TrimSpace(text) == "" {
		err += "comment cannot be blank"
	}

	if len(err) == 0 {
		return nil
	}

	return errors.New(err)
}

func ValidateTitle(titleInterface interface{}) error {
	title := titleInterface.(string)
	var err string

	if strings.TrimSpace(title) == "" {
		err += "title cannot be blank"
	}

	if len(err) == 0 {
		return nil
	}

	return errors.New(err)
}

func ValidateEmail(emailInterface interface{}) error {
	email := emailInterface.(string)
	var err string

	exp := regexp.MustCompile(`(?m)^[A-Za-z0-9._%+-]+@(?:[A-Za-z40-9-]+\.)+[A-Za-z]{2,10}$`)

	if !exp.Match([]byte(email)) {
		err += "email is not valid"
	}
	if strings.TrimSpace(email) == "" {
		err += "email cannot be blank"
	}

	if len(err) == 0 {
		return nil
	}

	return errors.New(err)
}

func ValidateNickname(nicknameInterface interface{}) error {
	nickname := nicknameInterface.(string)
	var err string

	exp := regexp.MustCompile(`(?m)[a-z0-9A-Z_-]]`)

	if exp.Match([]byte(nickname)) {
		err += "can contain only low and upper case letters and numbers"
	}

	if strings.TrimSpace(nickname) == "" {
		err += "nickname cannot be blank"
	}

	if len(err) == 0 {
		return nil
	}

	return errors.New(err)
}

func ValidatePassword(passwordInterface interface{}) error {
	password := passwordInterface.(string)
	var err string

	expCapLetter := regexp.MustCompile(`(?m)[A-Z]`)
	expLowCaseLetter := regexp.MustCompile(`(?m)[a-z]`)


	if !(len(password) >= 8 && len(password) <= 24) {
		err += "length of the password has to be more then 8 and less then 24 "
	}
	if !(expCapLetter.Match([]byte(password))) {
		err += "must contain capital letter "
	}
	if !(expLowCaseLetter.Match([]byte(password))) {
		err += "must contain low case letter "
	}
	if strings.TrimSpace(password) == "" {
		err += "password cannot be blank"
	}

	if len(err) == 0 {
		return nil
	}

	return errors.New(err)
}