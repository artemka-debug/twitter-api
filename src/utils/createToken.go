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
			ExpirationTime: jwt.NumericDate(now.Add(24 * 30 * 12 * time.Hour)),
			NotBefore:      jwt.NumericDate(now.Add(30 * time.Minute)),
			IssuedAt:       jwt.NumericDate(now),
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
