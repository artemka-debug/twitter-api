package api

import (
	"encoding/json"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func SubscribeToNotifications(c *gin.Context) {
	body := c.Keys["body"].(map[string]interface{})["body"].(map[string]interface{})
	sub := body["sub"].(string)

	var subscription webpush.Subscription
	_ = json.Unmarshal([]byte(sub), &subscription)

	var subscribed int

	_ = db.DB.QueryRow(`select id from subscription where userId = ?`, body["userId"]).Scan(&subscribed)

	if subscribed != 0 {
		_, err := db.DB.Exec(`update subscription set 
                        userId = ?, 
                        endpoint = ?, 
                        auth = ?, 
                        p256dh = ?
						where userId = ?`,
			body["userId"], subscription.Endpoint, subscription.Keys.Auth, subscription.Keys.P256dh, body["userId"])

		if err != nil {
			utils.HandleError([]string{"could not update your subscription"}, err.Error(), c, 500)
		}

		utils.SendPosRes(c, 200, gin.H{})
		return
	}

	_, err := db.DB.Exec(`insert into subscription(userId, endpoint, auth, p256dh) 
								values (?, ?, ?, ?)`,
		body["userId"],
		subscription.Endpoint,
		subscription.Keys.Auth,
		subscription.Keys.P256dh)

	if err != nil {
		utils.HandleError([]string{"could not subscribe to notification"}, err.Error(), c, 500)
		return
	}

	utils.SendPosRes(c, 201, gin.H{})
}
