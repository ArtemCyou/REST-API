package controller

import (
	"example/REST-API-APREL/service"
	"github.com/gin-gonic/gin"
)

type credentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService: jwtService,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}

	isAuthenticated := c.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		//return c.jWtService.GenereteToken(credentials.username, true)
		return c.jwtService.GenerateToken(credentials.Username,credentials.Password)
	}
	return ""

}
