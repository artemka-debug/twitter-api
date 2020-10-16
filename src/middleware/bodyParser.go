package middleware

import (
	"encoding/json"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"reflect"
)

func BodyParser(c *gin.Context) {
	data, _ := ioutil.ReadAll(c.Request.Body)

	if len(data) == 0 {
		utils.HandleError([]string{"no data was send"}, "req body is empty", c, 400)
		c.Abort()
		return
	}

	t := reflect.TypeOf(utils.GetStructByUri()[c.Request.URL.Path])
	v := reflect.New(t.Elem())
	newP := v.Interface()
	err := json.Unmarshal(data, newP)

	if err != nil {
		utils.HandleError([]string{"try again"}, err.Error(), c, 400)
		return
	}

	c.Set("body", newP)
	c.Next()
}
