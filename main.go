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
	gin.SetMode(gin.ReleaseMode)
	PORT, exists := os.LookupEnv("PORT")

	if !exists {
		PORT = "8000"
	}
	fmt.Println("PORT", PORT)
	r.Use(middleware.SetHeaders)

	r.GET("/me", api.Me)
	r.POST("/sign-up", middleware.BodyParser, middleware.InputValidate, api.SignUp)
	r.POST("/login", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.Login)
	r.PUT("/reset-password", middleware.BodyParser, middleware.InputValidate, api.ResetPassword)
	r.DELETE("/remove-user", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.RemoveUser)
	r.POST("/post", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.Post)

	err := r.Run(fmt.Sprintf(":%s", PORT))

	if err != nil {
		fmt.Print("Failed listen", err)
	}
}
