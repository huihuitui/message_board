package dao

import "message-board/model"

func Comment(m model.Message) (err error) {
	_, err = DB.Exec("insert into message (send_name,rec_name,detail) values (?,?,?)", m.SendName, m.RecUName, m.Detail)
	return
}
