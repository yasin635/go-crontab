package config

import (
	"os"
	"encoding/json"
	"fmt"
)

type ConfigJson struct {
	Mysql_host string //mysql host
	Mysql_port string
	Mysql_user string
	Mysql_pass string
	Mysql_dbname string
}

var Config = &ConfigJson{}

func LoadConfig(path string) error {
	file,_ := os.Open(path + "/crontab/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Config)
	if err != nil {
		fmt.Println("读取json文件失败")
	}
	return err
}

//func Get() map[string]string  {
//	currentPath, _ := os.Getwd()
//	file,_ := os.Open(currentPath +"/crontab/config.json")
//	defer file.Close()
//	conf := new(ConfigMysql)
//	decoder := json.NewDecoder(file)
//	err := decoder.Decode(&conf)
//	if err != nil {
//		//return err
//		fmt.Println("读取json文件读取不要")
//
//	}
//	//fmt.Println(conf.Host)
//	return conf
//}
func init()  {
	//currentPath, _ := os.Getwd()
	//file,_ := os.Open(currentPath + "/crontab/config/config.json")
	//defer file.Close()
	//conf := new(ConfigMysql)
	//decoder := json.NewDecoder(file)
	//err := decoder.Decode(&conf)
	//if err != nil {
	//	fmt.Println("读取json文件读取不要")
	//}
	//fmt.Println(conf)
}