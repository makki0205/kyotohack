package service

import "github.com/makki0205/kyotohack/model"

var users []model.User
var id = 0

func init() {
	StoreUser(model.User{
		Email:    "hoge",
		Password: "foo",
		NicName:  "makki",
	})
	StoreUser(model.User{
		Email:    "hoge",
		Password: "foo",
		NicName:  "jin",
	})
}
func StoreUser(user model.User) {
	user.ID = userID()
	users = append(users, user)
}
func userID() int {
	id++
	return id
}

func GetByID(id int) *model.User {
	for _, u := range users {
		if u.ID == id {
			return &u
		}
	}
	return nil
}
func GetUsers() []model.User {
	return users
}
