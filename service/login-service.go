package service

import (
	"example/REST-API-APREL/cntl"
	"example/REST-API-APREL/entity"
	"gorm.io/gorm"
)

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	authorizationUsername string
	authorizationPassword string
}

func NewLoginService() LoginService {
	return &loginService{}
	//&loginService{
	//	authorizationUsername: "cherbis",
	//	authorizationPassword: "1234",
	//}
}

func (l *loginService) Login(username, password string) bool {
	user := entity.User{}
	conn := cntl.UserRepository
	db := conn.GetDB()
	err := db.Table("users").Where("login = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false
		}
	}
	l.authorizationUsername = user.Login
	l.authorizationPassword = user.Password

	return l.authorizationUsername == username &&
		l.authorizationPassword == password
}

//
//func GetDB() *gorm.DB {
//return db
//}
