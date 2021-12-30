package entity

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	SuccessAuth = iota
	IncorrectPass
	Block
)

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Login    string `json:"login" binding:"required" gorm:"type:varchar(32);UNIQUE"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(32)"`
	BadEntry int8   `json:"bad_entry" binding:"required,gte=1,lte=5"`
}

type Tokens struct {
	Token   string `json:"token" gorm:"type:varchar(256)"`
	Visitor User   `json:"visitor" gorm:"foreignkey:UserID"`
	UserID  uint64 `json:"-"`
}

type AuditAuthorize struct {
	Visitor   User      `json:"visitor" gorm:"foreignkey:UserID"`
	UserID    uint64    `json:"-"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	Event     int8      `json:"event"`
}

type JwtCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}
