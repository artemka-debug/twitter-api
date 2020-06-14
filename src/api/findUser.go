package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"net/url"
)

func FindUser(c *gin.Context) {
	test, errorParsing := url.ParseQuery(c.Request.URL.String())

	if errorParsing != nil {
		utils.HandleError([]string{"could not parse query"}, errorParsing.Error(), c, 500)
	}

	if _, ok := test["/users?nickname"]; !ok {
		utils.HandleError([]string{"no query params were provided"}, "no query params were provided", c, 500)
		return
	}

	if len(test["/users?nickname"]) == 0 {
		utils.HandleError([]string{"nickname is empty"}, "nickname is empty", c, 400)
		return
	}

	rows, errorFinding := db.DB.Query(`select id, nickname, status from users where nickname like concat('%', ?, '%')`, test["/users?nickname"][0])
	defer rows.Close()

	if errorFinding != nil {
		utils.HandleError([]string{"error finding users"}, errorFinding.Error(), c, 500)
		return
	}

	var users []utils.Users

	for rows.Next() {
		var user utils.Users

		if err := rows.Scan(&user.Id, &user.Nickname, &user.Status); err != nil {
			utils.HandleError([]string{"could not get all users"}, err.Error(), c, 500)
			return
		}

		users = append(users, user)
	}

	utils.SendPosRes(c, 200, gin.H{
		"users": users,
	})
}
