package utils

import (
	"github.com/gbrlsnchs/jwt/v3"
)

type CommentSchema struct {
	PostId int    `json:"post_id"`
	Text   string `json:"text"`
	UserId int    `json:"user_id"`
}

type CustomPayload struct {
	jwt.Payload
	Id       int
	Password string
}

type PostSchema struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	UserId int    `json:"user_id"`
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
