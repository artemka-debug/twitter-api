package utils

import (
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/env"
)

func SendNotification(userId int, msg string) {
	var email, endpoint, auth, p256dh string

	errorGettingSubscription := db.DB.QueryRow(`select email, endpoint, auth, p256dh from subscription inner join users u on subscription.userId = u.id where userId = ?`, userId).Scan(&email, &endpoint, &auth, &p256dh)

	fmt.Println("errorGettingSubscription", errorGettingSubscription)

	subscription := webpush.Subscription{
		Endpoint: endpoint,
		Keys:     webpush.Keys{
			Auth:   auth,
			P256dh: p256dh,
		},
	}

	res, errorSendingNotification := webpush.SendNotification([]byte(msg), &subscription, &webpush.Options{
		Subscriber:      email,
		VAPIDPublicKey:  env.VapidPublicKey,
		VAPIDPrivateKey: env.VapidPrivateKey,
		TTL:             60,
	})

	if errorSendingNotification != nil {
		fmt.Println("errorSendingNotification", errorSendingNotification)
	} else {
		fmt.Println("SENT", res)
	}
}
