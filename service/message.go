package service

import (
	"message-board/dao"
	"message-board/model"
)

func Get() (m []model.Message, err error) {
	m, err = dao.Get()
	return
}
func GetComment(s string) (m []model.Message, err error) {
	m, err = dao.GetComment(s)
	return
}
func GetUserMessage(s string) (m []model.Message, err error) {
	m, err = dao.GetUserMessage(s)
	return
}
func ModifyMessage(user string, newDetail string, detail string) (err error) {
	err = dao.ModifyMessage(user, newDetail, detail)
	return
}
func Comment(m model.Message) (err error) {
	err = dao.Comment(m)
	return
}
func Delete(s string) (err error) {
	err = dao.Delete(s)
	return
}
