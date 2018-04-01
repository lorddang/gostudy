package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {


	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("error")
		return
	}
	rows, err := db.Query("SELECT * FROM  test" )
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		fmt.Println(rows)
	}
}
