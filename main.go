package main

import (
	"github.com/gin-gonic/gin"
	"github.com/makki0205/kyotohack/controller"
)

func main() {
	r := gin.Default()
	r.GET("/user", controller.GetAllUser)
	r.POST("/user", controller.CreateUser)
	r.Run(":8000")
}
