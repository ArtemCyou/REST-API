package service

type LoginService interface {
	Login(username, password string) bool
}

type loginService struct {
	authorizationUsername string
	authorizationPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizationUsername: "cherbis",
		authorizationPassword: "1234",
	}
}

func (l *loginService) Login(username, password string) bool {
	//user := &loginService{}
	//db:= repository.NewUserRepository()
	//err := db.c Table("users").Where("login = ?", login).First(user).Error
	//if err != nil {
	//	if err == gorm.ErrRecordNotFound {
	//		return false
	//	}
	//}
	return l.authorizationUsername == username &&
		l.authorizationPassword == password
}
//
//func GetDB() *gorm.DB {
//return db
//}
