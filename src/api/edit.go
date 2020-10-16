package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func Edit(c *gin.Context) {
	body := c.Keys["body"].(*utils.EditSchema)
	userId := c.Keys["userId"].(int)

	if body.Nickname == "" {
		_, errorUpdating := db.DB.Exec(`update Users
                                          set status = ?
                                          where Users.id = ?`, body.Status, userId)
		if errorUpdating != nil {
			utils.HandleError([]string{"cannot edit your profile"}, errorUpdating.Error(), c, 500)
			return
		}
		utils.SendPosRes(c, 204, gin.H{})
		return
	}

	_, errorUpdating := db.DB.Exec(`update Posts, Users
                    left join Posts P on Users.id = P.user_id
    				left join comments C on users.id = C.user_id
                                          set P.nickname = ?,
                                              C.nickname = ?,
                                              C.nickname = ?,
                                              users.nickname = ?,
                                              status = ?
                                          where Users.id = ?`, body.Nickname, body.Nickname, body.Nickname, body.Nickname, body.Status, userId)
	if errorUpdating != nil {
		utils.HandleError([]string{"cannot edit your profile"}, errorUpdating.Error(), c, 500)
		return
	}

	utils.SendPosRes(c, 200, gin.H{})
}
