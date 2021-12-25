package controller

import (
	"example/REST-API-APREL/entity"
	"example/REST-API-APREL/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	FindAll() []entity.User
	Create(ctx *gin.Context) error
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.User {
	return c.service.FindAll()
}

func (c *controller) Create(ctx *gin.Context) error {
	var user entity.User
	//err := ctx.BindJSON(&user)
	if err := ctx.ShouldBindJSON(&user); err != nil {
	 	return err
	 	}
	c.service.Create(user)
	return nil

}
