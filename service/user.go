package service

import (
	"message-board/dao"
	"message-board/model"
	"message-board/util"
)

func CreateUser(u model.User) (err error) {
	u.Password, err = util.Encrypt(u.Password)
	err = dao.InsertUser(u)
	return
}
func SearchUserByName(name string) (u model.User, err error) {
	u, err = dao.SearchUserByName(name)
	return
}
