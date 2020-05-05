package utils

import (
	"github.com/gin-gonic/gin"
)

func SendErrorRes(c *gin.Context, err, token string, code int) {
	c.JSON(code, Res{
		Token: token,
		Result: false,
		Error: err,
		Id: 0,
	})
}

func SendPosRes(token string, c *gin.Context, code int, id int) {
	c.JSON(code, Res{
		Token: token,
		Result: true,
		Error: "",
		Id: id,
	})
}
