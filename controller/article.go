package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/makki0205/kyotohack/service"
)

func GetArticle(c *gin.Context) {
	var res []Article

	for _, value := range service.GetArticle() {
		user := service.GetByID(value.UserID)
		res = append(res, Article{
			Title: value.Title,
			Body:  value.Body,
			User: User{
				Email:   user.Email,
				NicName: user.NicName,
			},
		})
	}
}

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	User  User   `json:"user"`
}

type User struct {
	Email   string `json:"email"`
	NicName string `json:"nick_name" binding:"required"`
}
