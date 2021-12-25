package service

import "example/REST-API-APREL/entity"

type UserService interface {
	FindAll()[]entity.User
	Create(user entity.User) entity.User

}

type userService struct {
	users []entity.User
}

func New() UserService {
	return &userService{}
}

func (u *userService)Create(user entity.User)entity.User  {
u.users = append(u.users, user)
return user
}

func (u *userService)FindAll() []entity.User  {
 return u.users
}