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
	gin.SetMode(gin.DebugMode)
	PORT, exists := os.LookupEnv("PORT")

	if !exists {
		PORT = "8000"
	}
	fmt.Println("PORT", PORT)
	r.Use(middleware.SetHeaders)
	r.Use(middleware.BodyParser)

	r.POST("/me", api.Me)
	r.POST("/sign-up", api.SignUp)
	r.POST("/login", api.Login)
	r.PUT("/reset-password", api.ResetPassword)

	err := r.Run(fmt.Sprintf(":%s", PORT))

	if err != nil {
		fmt.Print("Failed listen", err)
	}
}
