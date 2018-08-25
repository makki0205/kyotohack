package service

import (
	"fmt"
	"testing"

	"github.com/makki0205/kyotohack/model"
)

func TestGetUsers(t *testing.T) {
	StoreUser(model.User{
		Email:    "hoge@hoge.com",
		Password: "hogehoge",
		NicName:  "hoge",
	})
	StoreUser(model.User{
		Email:    "hoge2@hoge.com",
		Password: "hogehoge",
		NicName:  "hoge2",
	})
	fmt.Println(GetUsers())
}
