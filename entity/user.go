package entity

type User struct {
	//ID       int    `json:"id" binding:"serial primary key"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	BadEntry int    `json:"bad_entry" binding:"required"`
}
