package utils

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/secret"
	"github.com/gbrlsnchs/jwt/v3"
	"time"
)

func CreateToken(id int) string {
	now := time.Now()
	pl := CustomPayload{
		Payload: jwt.Payload{
			IssuedAt: jwt.NumericDate(now),
		},
		Id: id,
	}

	token, err := jwt.Sign(pl, secret.AppKey)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(token)
}
