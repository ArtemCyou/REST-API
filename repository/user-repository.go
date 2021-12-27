package repository

import (
	"example/REST-API-APREL/entity"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type UserRepository interface {
	Creates(user entity.User)
	Delete(user entity.User)
	FindAll() []entity.User
	CloseDB()
}

type database struct {
	connection *gorm.DB
}


func NewUserRepository() UserRepository {
	db, err := gorm.Open(sqlite.Open("sqlLite.db"), &gorm.Config{})
	if err != nil {
		panic("failed connection database")
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("failed migrate database")
	}

	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
sql, err:= db.connection.DB()
	if err != nil {
		panic("failed connection for close database")
	}
	err = sql.Close()
	if err != nil {
		panic("failed close database")
	}
}

func (db *database) Creates(user entity.User) {
	db.connection.Create(&user)
}

func (db *database) Delete(user entity.User) {
	db.connection.Delete(&user)
}

func (db *database) FindAll() []entity.User {
	var users []entity.User
	db.connection.Find(&users)
	return users
}
