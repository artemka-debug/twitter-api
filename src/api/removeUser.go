package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func RemoveUser(c *gin.Context) {
	userId := c.Keys["userId"].(int)
	var password string
	errorSelectingFromDb := db.DB.QueryRow(`select password from users where id = ?`, userId).Scan(&password)

	if errorSelectingFromDb != nil {
		utils.HandleError([]string{"you dont have an account, you need to sign up"}, errorSelectingFromDb.Error(), c, 400)
		return
	}
	_, errorDeletingUser := db.DB.Query(`delete users from users
												  left join posts p on users.id = p.user_id
												  left join comments c on users.id = c.user_id
												  left join liked_posts l on users.id = l.user_id
												where id = ?`, userId)

	if errorDeletingUser != nil {
		utils.HandleError([]string{"could not delete user, try again"}, errorDeletingUser.Error(), c,  403)
	}

	utils.SendPosRes(c, 200, gin.H{})
}
