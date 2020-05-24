package parseBody

import (
	"encoding/json"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func BodySubscription(data []byte, c *gin.Context) {
	var body utils.Subscription

	errorDecoding := json.Unmarshal(data, &body)
	defer c.Request.Body.Close()

	if errorDecoding != nil {
		utils.HandleError([]string{"could not parse your data, try again"}, errorDecoding.Error(), c, 400)
		c.Abort()
	}

	c.Set("body", body)
}
