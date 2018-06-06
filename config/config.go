package config

import (
	"os"
	"encoding/json"
	"fmt"
)

//import "fmt"

type ConfigMysql struct {
		Host string
		Port int32
		User string
		Pass string
}

//var conf = ConfigMysql{Host:"host",Port:3306,User:"root",Pass:"123456"}
func LoadConfig()  {


	file,_ := os.Open("crontab/config/config.json")
	defer file.Close()
	conf := ConfigMysql{}
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("读取json文件读取不要")
	}
	fmt.Println(conf.Host)
}
func init()  {

}