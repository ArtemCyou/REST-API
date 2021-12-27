package entity

type User struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Login    string `json:"login" binding:"required" gorm:"type:varchar(32)"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(32)"`
	BadEntry int8    `json:"bad_entry" binding:"required,gte=1,gte=5"`
}
