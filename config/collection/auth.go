package collection

import (
	v1 "golang-jwt/app/controllers/v1"
	"golang-jwt/config/middleware"
	"golang-jwt/lib/env"

	"github.com/gin-gonic/gin"
)

func User(main *gin.RouterGroup) {
	authController := v1.InitAuthController()
	group := main.Group(env.String("InternalRouting.V1.Prefix", ""))
	group.POST(env.String("InternalRouting.V1.GetUser.Send", "") , authController.RegisterUser)
	group.POST(env.String("InternalRouting.V1.LoginUser.Send", "") , authController.LoginUser)
	
	group.Use(middleware.JwtAuthMiddleware()) 
	{
		group.GET(env.String("InternalRouting.V1.GetUser.Send", "") , authController.GetUsers)
	}
}