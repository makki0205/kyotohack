package model

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	NicName  string `json:"nick_name" binding:"required"`
}

type Users []User
