package controller

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_latihan_tools")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TidakDipake() {
	//Ini cuma buat mysql bisa diimport
	x := mysql.ErrOldPassword
	log.Print(x)
}
