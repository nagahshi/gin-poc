package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nagahshi/gin-poc/service"
)

type AuthController interface {
	Login(ctx *gin.Context) string
}

type authController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

type Credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func NewAuthController(loginService service.LoginService, jwtService service.JWTService) AuthController {
	return &authController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (auth *authController) Login(ctx *gin.Context) string {
	var credentials Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := auth.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return auth.jwtService.GenerateToken(credentials.Username, true)
	}

	return ""
}
