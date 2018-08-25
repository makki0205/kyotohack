package service

import "github.com/makki0205/kyotohack/model"

var users []model.User

func StoreUser(user model.User) {
	users = append(users, user)
}

func GetUsers() []model.User {
	return users
}
