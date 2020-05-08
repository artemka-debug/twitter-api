package api

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	body := c.Keys["body"].(utils.EditSchema)
	fmt.Println("USER", body)

	if body.Nickname == "" {
		utils.HandleError([]string{"nickname is empty"}, "fuck you stupid, just add valid nickname)", c, 400)
		return
	}
	userId := c.Keys["userId"].(int)
	_, errorUpdating := db.DB.Exec(`update Posts, Users
                    left join Posts P on Users.id = P.user_id
    				left join comments C on users.id = C.user_id
                                          set P.nickname = ?,
                                              C.nickname = ?,
                                              C.nickname = ?,
                                              status = ?
                                          where Users.id = ?`, body.Nickname, body.Nickname, body.Nickname, body.Status, userId)
	if errorUpdating != nil {
		utils.HandleError([]string{"cannot edit your profile"}, errorUpdating.Error(), c, 500)
		return
	}

	utils.SendPosRes(c, 200, gin.H{})
}
