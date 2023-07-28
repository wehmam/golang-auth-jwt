package config

import (
	"golang-jwt/config/collection"
	"golang-jwt/db/models"
	"golang-jwt/lib/env"

	"github.com/gin-gonic/gin"
)

var (
	Routers = gin.Default()
)

func init() {
	models.Connect()
	v1 := Routers.Group(env.String("InternalRouting.Base", ""))
	collection.User(v1)
}