package api

import (
	"github.com/artemka-debug/twitter-api/src/db"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func Post(c *gin.Context) {
	body := c.Keys["body"].(utils.PostSchema)
	id := -1
	errorSelectingFromDb := db.DB.QueryRow(`select User_PK from users where User_PK = ?`, body.Id).Scan(&id)

	if errorSelectingFromDb != nil  {
		utils.HandleError("you dont have an account, you need to sign up", c, 400)
		return
	}
	_, errorPosting := db.DB.Query(`insert into Posts(title, text, author, time, likes, user_id)
                                  values (?, ?, ?, ?, ?, ?)`, body.Title, body.Text, "", time.Now(), 0, id)

	if errorPosting != nil {
		utils.HandleError("error posting, try again", c, 500)
		return
	}

	utils.SendPosRes(strings.Split(c.GetHeader("Authorization"), " ")[1], c, 200, id)
}
