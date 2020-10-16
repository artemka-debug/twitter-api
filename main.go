package main

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/api"
	"github.com/artemka-debug/twitter-api/src/env"
	"github.com/artemka-debug/twitter-api/src/middleware"
	"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.SetHeaders)

	r.POST("/main", middleware.BodyParser, func(c *gin.Context) {
		utils.SendPosRes(c, 200, gin.H{
			"cool": "hi",
		})
	})
	r.POST("/google/auth", api.GoogleAuth)
	r.GET("/me", middleware.VerifyToken, api.Me)
	r.POST("/notification/subscribe", middleware.BodyParser, middleware.VerifyToken, api.SubscribeToNotifications)

	//// REGISTRATION
	r.POST("/sign-up", middleware.BodyParser, middleware.InputValidate, api.SignUp)
	r.POST("/login", middleware.BodyParser, middleware.InputValidate, api.Login)

	//// USER
	r.DELETE("/user", middleware.InputValidate, middleware.VerifyToken, api.RemoveUser)
	r.PUT("/user", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.Edit)
	r.PUT("/user/password", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.ChangePassword)
	r.PUT("/user/password/reset", middleware.BodyParser, middleware.InputValidate, api.ResetPassword)
	r.GET("/user/:id", api.GetUser)
	r.GET("/users", api.FindUser)

	//// TWEET
	r.DELETE("/tweet/:id", middleware.InputValidate, middleware.VerifyToken, api.RemovePost)
	r.POST("/tweet", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.Post)
	r.GET("/tweet/:id", api.GetPost)
	r.GET("/tweets", api.GetPosts)
	r.GET("/tweets/count", api.GetPostsCount)
	r.GET("/tweets/user/:id", api.GetPostsByUser)
	r.PUT("/tweet/like/:id", middleware.InputValidate, middleware.VerifyToken, api.LikeUnlikePost)

	//// COMMENT
	r.POST("/comment", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.AddComment)

	errorListening := r.Run(fmt.Sprintf(":%s", env.PORT))
	//http.DefaultTransport

	if errorListening != nil {
		fmt.Print("Failed listen", errorListening)
	}
}
