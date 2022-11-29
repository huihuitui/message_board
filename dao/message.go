package dao

import "message-board/model"

func Comment(m model.Message) (err error) {
	_, err = DB.Exec("insert into message (send_name,rec_name,detail,time) values (?,?,?,?)", m.SendName, m.RecUName, m.Detail, m.Time)
	return
}
func Delete(s string) (err error) {
	_, err = DB.Exec("update message set is_delete = 1 where send_name =?", s)
	return
}
