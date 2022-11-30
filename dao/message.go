package dao

import (
	"log"
	"message-board/model"
)

func Get() (m []model.Message, err error) {
	i := 0
	m = make([]model.Message, 1)
	Rows, err := DB.Query("select mid,uid,send_name,rec_name,detail,time,is_delete from message")
	if err != nil {
		log.Println(err)
		return
	}
	defer Rows.Close()
	for Rows.Next() {
		err = Rows.Scan(&m[i].MID, &m[i].UID, &m[i].SendName, &m[i].RecUName, &m[i].Detail, &m[i].Time, &m[i].IsDelete)
		if err != nil {
			log.Println(err)
			return
		}
		i++
	}
	return
}
func GetComment(s string) (m []model.Message, err error) {
	i := 0
	m = make([]model.Message, 1)
	Rows, err := DB.Query("select mid,uid,send_name,rec_name,detail,time,is_delete from message where rec_name =?", s)
	if err != nil {
		log.Println(err)
		return
	}
	defer Rows.Close()
	for Rows.Next() {
		err = Rows.Scan(&m[i].MID, &m[i].UID, &m[i].SendName, &m[i].RecUName, &m[i].Detail, &m[i].Time, &m[i].IsDelete)
		if err != nil {
			log.Println(err)
			return
		}
		i++
	}
	return
}
func GetUserMessage(s string) (m []model.Message, err error) {
	i := 0
	m = make([]model.Message, 1)
	Rows, err := DB.Query("select mid,uid,send_name,rec_name,detail,time,is_delete from message where send_name =?", s)
	if err != nil {
		log.Println(err)
		return
	}
	defer Rows.Close()
	for Rows.Next() {
		err = Rows.Scan(&m[i].MID, &m[i].UID, &m[i].SendName, &m[i].RecUName, &m[i].Detail, &m[i].Time, &m[i].IsDelete)
		if err != nil {
			log.Println(err)
			return
		}
		i++
	}
	return
}
func ModifyMessage(user string, newDetail string, detail string) (err error) {
	_, err = DB.Exec("update message set detail = ? where send_name =? and detail =?", newDetail, user, detail)
	return
}
func Comment(m model.Message) (err error) {
	_, err = DB.Exec("insert into message (uid,send_name,rec_name,detail,time) values (?,?,?,?,?)", m.UID, m.SendName, m.RecUName, m.Detail, m.Time)
	return
}
func Delete(s string) (err error) {
	_, err = DB.Exec("update message set is_delete = 1 where send_name =?", s)
	return
}
