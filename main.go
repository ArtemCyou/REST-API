package main

import (
	"example/REST-API-APREL/controller"
	"example/REST-API-APREL/middlewares"
	"example/REST-API-APREL/repository"
	"example/REST-API-APREL/service"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	//"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)

var (
	userRepository  repository.UserRepository  = repository.NewUserRepository()
	userService     service.UserService        = service.New(userRepository)
	userController  controller.UserController  = controller.New(userService)
	jwtService      service.JWTService         = service.NewJWTService()
	loginService    service.LoginService       = service.NewLoginService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

//авторизация. принимает логин, пароль отдает токен

//func getToken(c *gin.Context) {
//	var logJson struct {
//		Name     string `json:"name"`
//		Password string `json:"password"`
//	}
//	if err := c.ShouldBindJSON(&logJson); err != nil {
//		c.JSON(http.StatusUnauthorized, gin.H{"error:": err.Error()})
//		return
//	}
//	for _, v := range users {
//		if v.BadEntry >= 5 {
//			//заблокировать юзера
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "вы заблокированы"})
//			return
//		}
//		if logJson.Name == v.Login && logJson.Password == v.Password {
//			//сгенерировать токен
//			//сохранить токен в бд сессии
//			return
//		} else if logJson.Name == v.Login && logJson.Password != v.Password {
//			v.BadEntry += 1
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "неправильный пароль"})
//			return
//		}
//	}
//	c.JSON(http.StatusOK, gin.H{"status": "вы успешно авторизовались"})
//}
//удалить истории авторизации. check token in bd
func cleanAudit(c *gin.Context) {
	// get token
	//token := c.Param("token")

	//check token validation in sql
	//for i, t := range sessions {
	//	if t.Token == token {
	//		//по токену находим нужную сессию пользователя
	//		//clean audit authorization users
	//		sessions = append(sessions[:i], sessions[i+1:]...)
	//		c.IndentedJSON(http.StatusNoContent, t)
	//		return
	//	}
	//}
	c.IndentedJSON(http.StatusNotFound, gin.H{"сообщение": "такого токена не существует"})

}

func setupLogOutput() {
	f, _ := os.Create("ginOUT.txt")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	defer userRepository.CloseDB()
	setupLogOutput()

	router := gin.New()
	router.Use(gin.Recovery(), middlewares.Logger())
	//middlewares.BasicAuth()

	// Login Endpoint: Authentication + Token creation
	router.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := router.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "good auth user"})
		})
	}

	router.GET("/users", func(c *gin.Context) {
		//c.JSON(200,cntl.UserController.FindAll())
		c.JSON(200, userController.FindAll())
	})

	router.POST("/users", func(c *gin.Context) {
		//err:=cntl.UserController.Create(c)
		err := userController.Create(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status:": "User input is valid!"})
		}
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		//err:=cntl.UserController.Delete(c)
		err := userController.Delete(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status:": "Users delete!"})
		}
	})
	router.Run("localhost: 8080")

}
