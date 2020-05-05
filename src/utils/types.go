package utils

import (
	"github.com/gbrlsnchs/jwt/v3"
)

type Res struct {
	Token  string `json:"token"`
	Result bool   `json:"result"`
	Error  string `json:"error"`
	Id     int    `json:"id"`
}

type CustomPayload struct {
	jwt.Payload
	Email    string
	Password string
}

type PostSchema struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Id    int    `json:"id"`
}

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordSchema struct {
	Email string `json:"email"`
}

type RemoveUserSchema struct {
	Id       int    `json:"id"`
	Password string `json:"password"`
}

type SignupSchema struct {
	LoginSchema
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}
