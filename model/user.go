package model

type User struct {
	ID       int
	Email    string `json:"email"`
	Password string `json:"password"`
	NicName  string `json:"nick_name" binding:"required"`
}
type Article struct {
	Title  string
	Body   string
	UserID int
}
type Users []User
