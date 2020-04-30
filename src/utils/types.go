package utils

import (
	"github.com/gbrlsnchs/jwt/v3"
)

type Token struct {
	Token string `json:"token"`
}

type Res struct {
	Token  string `json:"token"`
	Result bool   `json:"result"`
	Error  string `json:"error"`
}

type CustomPayload struct {
	jwt.Payload
	email string
	password string
}

type UserSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname" validate:"required,email,min=2,max=40"`
	Status   string `json:"status" validate:"required,password,min=2,max=150"`
}
