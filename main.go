package main

import (
	"fmt"
	"crontab/config"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
)

var db *sql.DB
func init()  {
	//获取当前目录
	currentPath, err := os.Getwd()
	if err  != nil {
		fmt.Println("获取目录失败")
	}
	//fmt.Println(currentPath)
	//加载配置文件
	config.LoadConfig(currentPath)
}
func main()  {
	getlist()
}

//查询数据
func getlist()  {
	//uid:pass@tcp(host:port)/dbname?charset=utf8&parseTime=true
	//查询数据
	Conf := config.Config

	//用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8&parseTime=true
	dbsource := Conf.Mysql_user+":"+Conf.Mysql_pass+"@tcp("+Conf.Mysql_host+":"+Conf.Mysql_port+")/"+Conf.Mysql_dbname+"?charset=utf8&parseTime=true"
	//fmt.Println(dbsource)
	db, err := sql.Open("mysql", dbsource)
	//db, err := sql.Open("mysql", "root:Aa123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("Link database error: %s\n", err)
	}
	rows ,err := db.Query("SELECT topic FROM test.mqtt_logs")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next()  {
		var topic string
		if err := rows.Scan(&topic); err != nil {
			log.Fatal(err)
		}
		fmt.Println(topic)
	}
}

//连接数据
//func Db() sql.DB {
//	//uid:pass@tcp(host:port)/dbname?charset=utf8&parseTime=true
//	//查询数据
//	Conf := config.Config
//
//	//用户名:密码@tcp(主机:端口)/数据库名称?charset=utf8&parseTime=true
//	dbsource := Conf.Mysql_user+":"+Conf.Mysql_pass+"@tcp("+Conf.Mysql_host+":"+Conf.Mysql_port+")/"+Conf.Mysql_dbname+"?charset=utf8&parseTime=true"
//	fmt.Println(dbsource)
//	db, err := sql.Open("mysql", dbsource)
//	//db, err := sql.Open("mysql", "root:Aa123456@tcp(127.0.0.1:3306)/test")
//	if err != nil {
//		log.Fatalf("Open database error: %s\n", err)
//	}
//	defer db.Close()
//	err = db.Ping()
//	if err != nil {
//		log.Fatalf("Link database error: %s\n", err)
//	}
//
//	return db
//}