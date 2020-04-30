package middleware

import (
	"encoding/json"
	"errors"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func BodyParser(c *gin.Context) {
		var m map[string]interface{}
		data, _ := ioutil.ReadAll(c.Request.Body)

		if len(data) == 0 {
			utils.HandleError(errors.New("no body"), c)
			c.Abort()
			return
		}
		var body map[string]interface{}
		errorDecoding := json.Unmarshal(data, &body)
		defer c.Request.Body.Close()

		if utils.HandleError(errorDecoding, c) {
			c.Abort()
			return
		}

		c.Set("body", m)
		c.Next()
}
