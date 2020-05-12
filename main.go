package main

import (
	"fmt"
	"github.com/artemka-debug/twitter-api/src/api"
	"github.com/artemka-debug/twitter-api/src/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	r := gin.Default()
	PORT, exists := os.LookupEnv("PORT")

	if !exists {
		PORT = "3000"
	}
	fmt.Println("PORT", PORT)
	r.Use(middleware.SetHeaders)

	r.GET("/me", api.Me)

	// REGISTRATION
	r.POST("/sign-up", middleware.BodyParser, middleware.InputValidate, api.SignUp)
	r.POST("/login", middleware.BodyParser, middleware.InputValidate, api.Login)

	// USER
	r.DELETE("/user", middleware.InputValidate, middleware.VerifyToken, api.RemoveUser)
	r.PUT("/user", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.Edit)
	r.PUT("/user/password", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.ChangePassword)
	r.PUT("/user/password/reset", middleware.BodyParser, middleware.InputValidate, api.ResetPassword)
	r.GET("/user/:id", middleware.InputValidate, middleware.VerifyToken, api.GetUser)

	// TWEET
	r.DELETE("/tweet/:id", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.RemovePost)
	r.POST("/tweet", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.Post)
	r.GET("/tweet/:id", middleware.InputValidate, middleware.VerifyToken, api.GetPost)
	r.GET("/tweets", middleware.InputValidate, middleware.VerifyToken, api.GetPosts)
	r.PUT("/tweet/like/:id",  middleware.InputValidate, middleware.VerifyToken, api.LikeUnlikePost)

	// COMMENT
	r.POST("/comment", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.AddComment)

	err := r.Run(fmt.Sprintf(":%s", PORT))

	if err != nil {
		fmt.Print("Failed listen", err)
	}
}
