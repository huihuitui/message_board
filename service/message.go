package service

import (
	"message-board/dao"
	"message-board/model"
)

func Comment(m model.Message) (err error) {
	err = dao.Comment(m)
	return
}
