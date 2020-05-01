package utils

import "github.com/gbrlsnchs/jwt/v3"

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

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPassword struct {
	Email    string `json:"email"`
}

type SignupSchema struct {
	LoginSchema
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}
