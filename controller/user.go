package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/makki0205/kyotohack/model"
	"github.com/makki0205/kyotohack/service"
)

func CreateUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	service.StoreUser(user)
	c.Status(http.StatusCreated)
}

func GetAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetUsers())
}
