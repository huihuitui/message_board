package dao

import (
	"message-board/model"
)

func InsertUser(u model.User) (err error) {
	_, err = DB.Exec("insert into user (name,password) values (?,?)", u.UserName, u.Password)
	return
}
func SearchUserByName(s string) (u model.User, err error) {
	Row := DB.QueryRow("select * from user where name = ?", s)
	if err = Row.Err(); Row.Err() != nil {
		return
	}
	err = Row.Scan(&u.ID, &u.UserName, &u.Password)
	return
}
