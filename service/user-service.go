package service

import (
	"example/REST-API-APREL/entity"
	"example/REST-API-APREL/repository"
)

type UserService interface {
	FindAll()[]entity.User
	Delete(user entity.User)
	Create(user entity.User) entity.User

}

type userService struct {
	userRepository repository.UserRepository
}

func New(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (u *userService)Create(user entity.User)entity.User  {
u.userRepository.Create(user)
return user
}

func (u *userService) Delete(user entity.User){
	u.userRepository.Delete(user)
}

func (u *userService)FindAll() []entity.User  {
 return u.userRepository.FindAll()
}