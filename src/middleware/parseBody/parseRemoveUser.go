package parseBody

import (
	"encoding/json"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func BodyRemoveUser(data []byte, c *gin.Context) {
	var body utils.RemoveUserSchema

	errorDecoding := json.Unmarshal(data, &body)
	defer c.Request.Body.Close()

	if errorDecoding != nil {
		utils.HandleError("could not parse your data, try again", c, 400)
		c.Abort()
	}

	c.Set("body", body)
}
