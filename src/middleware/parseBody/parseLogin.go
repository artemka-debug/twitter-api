package parseBody

import (
	"encoding/json"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func BodyLogin(data []byte, c *gin.Context) {
	var body utils.LoginSchema

	errorDecoding := json.Unmarshal(data, &body)
	defer c.Request.Body.Close()

	if errorDecoding != nil {
		utils.HandleError("could not parse your data, try again", c, 400)
		c.Abort()
	}

	c.Set("body", body)
}

