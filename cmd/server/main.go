package server

import (
	//"github.com/golang/glog"
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const errDB = "Fatal error db: %s \n"
var db *gorm.DB

type User struct {
	gorm.Model
	login string
	Password string
}

// orm initialization
func IntMigration()  {
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
if err != nil {
log.Fatalf(errDB, err)
}


db.AutoMigrate(&User{})
}

func main() {


	// run server
	defer glog.Flush()
	app.Run(&db.Database{ORM: orm})
}
