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
	return l.authorizationUsername == username &&
		l.authorizationPassword == password
}


