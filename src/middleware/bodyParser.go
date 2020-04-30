package middleware

import (
	"encoding/json"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func BodyParser(c *gin.Context) {
		var m map[string]interface{}
		decoder := json.NewDecoder(c.Request.Body)
		defer c.Request.Body.Close()

		errorDecoding := decoder.Decode(&m)
		if utils.HandleError(errorDecoding, c) {
			c.Abort()
			return
		}

		c.Set("body", m)
		c.Next()
}
