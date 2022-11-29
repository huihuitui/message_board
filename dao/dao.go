package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/message_board?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = db
	fmt.Println(DB.Ping())
}
