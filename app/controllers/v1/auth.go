package v1

import (
	"golang-jwt/db/models"
	"golang-jwt/lib/encrypt"
	"golang-jwt/lib/env"
	"golang-jwt/request"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	// "github.com/golang-jwt/jwt/v4"
)

type authController struct {
	Name    string
	BaseUrl string
}

func InitAuthController() *authController {
	baseUrl := env.String("MainSetup.ServerHost", "")
	return &authController{
		Name:    "AuthController - ",
		BaseUrl: baseUrl,
	}
}

func (ctrl *authController) GetUsers(c *gin.Context) {
	var users []models.User
	
	// models.DB.Raw("SELECT * FROM users").Scan(&users) // option 2 for get all users
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"status" : true,
		"data" : users,
	})
}

func (crtl *authController) RegisterUser(c *gin.Context) {
	var req request.ReqUsers
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : err.Error(),
		})
		return 
	}

	encrypt, err := encrypt.EncrpytFromString(req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error" : err.Error(),
		})
		return 
	}

	bodyParams := models.User{
		Username: req.Username,
		Password: encrypt,
	}

	errInsertDb := models.DB.Create(&bodyParams).Error
	if errInsertDb != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : false, 
			"error" : errInsertDb,
		})

		return
	}

	c.JSON(200, gin.H{
		"data" : req,
		"status" : true,
		"message" : "Berhasil Insert User",
	})
}

func (ctrl *authController) LoginUser(c *gin.Context) {
	var loginReq request.ReqUsers


	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : false,
			"Error" : err.Error(),
		})
		return 
	}

	bodyParams := models.User{
		Username: loginReq.Username,
		Password: loginReq.Password,
	}


	err := models.DB.Where("username= ?", bodyParams.Username).Take(&bodyParams).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : false,
			"Error" : err.Error(),
		})
		return
	}

	encrypt := encrypt.CompareHashPassword(bodyParams.Password, loginReq.Password)
	if !encrypt {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : false,
			"Error" : "Password Salah",
		})
		return
	}


	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = bodyParams.Id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	
	accessToken, err := token.SignedString([]byte("TESTING JWT AUTH WITH GOLANG"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status" : false,
			"Error" : err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" : true,
		"data" : gin.H{
			"accessToken" : accessToken,
			"expiredAt" : claims["exp"],
		},
	})
}
