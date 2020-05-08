package utils

import (
	"github.com/gbrlsnchs/jwt/v3"
)

type CommentSchema struct {
	PostId int    `json:"post_id"`
	Text   string `json:"text"`
}

type CustomPayload struct {
	jwt.Payload
	Id       int
}

type EditSchema struct {
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}

type PostSchema struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
}

type LoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Password string `json:"password"`
}

type ResetPasswordSchema struct {
	Email string `json:"email"`
}

type SignupSchema struct {
	LoginSchema
	Nickname string `json:"nickname"`
	Status   string `json:"status"`
}
