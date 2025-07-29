package controllers

import (
	"net/http"
	"time"

	"github.com/chonghua1013/ProductAnalysis/internal/models"
	"github.com/chonghua1013/ProductAnalysis/internal/services"
	"github.com/chonghua1013/ProductAnalysis/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthController struct {
	userService services.UserService
	jwtSecret   string
}

func NewAuthController(userService services.UserService, jwtSecret string) *AuthController {
	return &AuthController{
		userService: userService,
		jwtSecret:   jwtSecret,
	}
}

// Login 用户登录
func (ac *AuthController) Login(c *gin.Context) {
	var loginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// 在实际应用中，这里应该验证用户密码
	user, err := ac.userService.GetUserByEmail(loginReq.Email)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// 简化的密码验证 - 实际应该使用bcrypt等加密验证
	if user.Password != loginReq.Password {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	// 创建JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(ac.jwtSecret))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.RespondWithJSON(c, http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}

// Register 用户注册
func (ac *AuthController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// 检查用户是否已存在
	_, err := ac.userService.GetUserByEmail(user.Email)
	if err == nil {
		utils.RespondWithError(c, http.StatusConflict, "User already exists")
		return
	}

	// 创建用户
	newUser, err := ac.userService.CreateUser(&user)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithJSON(c, http.StatusCreated, newUser)
}
