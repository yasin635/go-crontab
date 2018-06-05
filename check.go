package main

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "reflect"
	"database/sql"
	"fmt"
)



func main(){
	db, err := sql.Open("mysql", "root=Aa123456@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
		println(err)
		println("1212")
	}
	db.Query("use test");
	fmt.Println("ss")
	query, err := db.Query("select * from mqtt_logs")
	fmt.Print(query)
}