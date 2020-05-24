package main

import (
	"github.com/SherClockHolmes/webpush-go"
	"github.com/artemka-debug/twitter-api/src/api"
	"github.com/artemka-debug/twitter-api/src/env"
	"github.com/artemka-debug/twitter-api/src/middleware"
	"github.com/artemka-debug/twitter-api/src/utils"

	//"github.com/artemka-debug/twitter-api/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	//"github.com/reactivex/rxgo/v2"

	"fmt"
	"net/http"
	"os"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wshandler(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		//utils.HandleError([]string{"bad connection"}, " ", c, 500)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("ERROR READING", err)
			return
		}

		fmt.Println(string(msg))
		errorMessaging := conn.WriteMessage(t, msg)

		if errorMessaging != nil {
			fmt.Println("ERROR SENDING", errorMessaging)
			return
		}
	}
}

func main() {
	r := gin.Default()
	PORT, exists := os.LookupEnv("PORT")

	if !exists {
		PORT = "3000"
	}
	fmt.Println("PORT", PORT)
	r.Use(middleware.SetHeaders)

	r.GET("/me", middleware.VerifyToken, api.Me)
	r.POST("/notification/subscribe", middleware.BodyParser, func(c *gin.Context) {
		body := c.Keys["body"].(utils.Subscription)
		subscription := webpush.Subscription{
			Endpoint: body.EndPoint,
			Keys:     body.Keys,
		}

		res, err := webpush.SendNotification([]byte("Test"), &subscription, &webpush.Options{
			Subscriber:      "dovgopoly000@gmail.com",
			VAPIDPublicKey:  env.VapidPublicKey,
			VAPIDPrivateKey: env.VapidPrivateKey,
			TTL:             60,
		})

		fmt.Println(res)


		if err != nil {
			fmt.Println(err)
		}
	})

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
	r.PUT("/tweet/like/:id", middleware.InputValidate, middleware.VerifyToken, api.LikeUnlikePost)

	// COMMENT
	r.POST("/comment", middleware.BodyParser, middleware.InputValidate, middleware.VerifyToken, api.AddComment)

	errorListening := r.Run(fmt.Sprintf(":%s", PORT))

	if errorListening != nil {
		fmt.Print("Failed listen", errorListening)
	}
}
