package main

import (
	"example/REST-API-APREL/middlewares"
	"example/REST-API-APREL/controller"
	"example/REST-API-APREL/service"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
	//"database/sql"
	_ "github.com/lib/pq"
	"net/http"
)
//user struct
//type user struct {
//	ID       int    `json:"id" binding:"serial primary key"`
//	Login    string `json:"login" binding:"required"`
//	Password string `json:"password" binding:"required"`
//	BadEntry int    `json:"bad_entry" binding:"required"`
//}

//var users = []user{
//	{Login: "root", Password: "4334384334", BadEntry: 0},
//}

//условная структура будет реализованна на jwt
type session struct {
	Token    *string   `json:"token"`
	TimeLive time.Time `json:"time_live"`
}

//var sessions = []session{
//	//{Token: "665464", TimeLive: '' },
//}

//аудит авторизации
type auditAuthruzation struct {
	Name  string      `json:"name"`
	Time  time.Time `json:"time"`
	Event string    `json:"event"`
}

var auth = []auditAuthruzation{
	//{Name: user{Login: "root"}, Time: '15', Event: "authorization"},
}
var (
	userService service.UserService = service.New()
	userController controller.UserController = controller.New(userService)
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

//история авторизации юзера. check token in bd. JSON: дата/время, событие.
func getHistAuth(c *gin.Context) {
	// get token
	//...
	//check token validation in sql
	//...
	//return
	var session auditAuthruzation
	c.IndentedJSON(http.StatusOK, session)
}

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

func setupLogOutput()  {
	f,_:= os.Create("ginOUT.txt")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
func main() {
	setupLogOutput()
	router := gin.New()
	router.Use(gin.Recovery(),middlewares.Logger(), middlewares.BasicAuth())

	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, userController.FindAll())
	})
	router.POST("/users", func(c *gin.Context) {
		err:= userController.Create(c)
		if err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status:": "User input is valid!"})
		}

	})

	//router.GET("/authorization", getToken)     //принимает логин и пароль
	router.GET("/users/:token", getHistAuth)   //json история авторизации
	router.DELETE("/users/:token", cleanAudit) //удаляет историю авторизации
	router.Run("localhost: 8080")

}
