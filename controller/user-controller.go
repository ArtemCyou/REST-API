package controller

import (
	"example/REST-API-APREL/entity"
	"example/REST-API-APREL/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController interface {
	FindAll() []entity.User
	Create(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
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

func (c *controller)Delete(ctx *gin.Context)error{
	var user entity.User
	id, err:= strconv.ParseUint(ctx.Param("id"),0,0)
	if err!=nil{
		return err
	}
	user.ID = id
	c.service.Delete(user)
	return nil
}
