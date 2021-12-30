package cntl

import (
	"example/REST-API-APREL/controller"
	"example/REST-API-APREL/repository"
	"example/REST-API-APREL/service"
)

var (
	UserRepository  repository.UserRepository  = repository.NewUserRepository()
	userService     service.UserService        = service.New(UserRepository)
	UserController  controller.UserController  = controller.New(userService)
	jwtService      service.JWTService         = service.NewJWTService()
	loginService    service.LoginService       = service.NewLoginService()
	LoginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)