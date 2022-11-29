package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreateUser(u model.User) (err error) {
	err = dao.InsertUser(u)
	return
}
func SearchUserByName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByName(name)
	return
}
