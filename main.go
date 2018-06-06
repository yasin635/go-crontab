package main

import (
	"fmt"
	"crontab/config"
	"os"
)

func init()  {
	//获取当前目录
	currentPath, err := os.Getwd()
	if err  != nil {
		fmt.Println("获取目录失败")
	}
	fmt.Println(currentPath)
	config.LoadConfig()
	fmt.Println(config.ConfigMysql{Host:currentPath})
}
func main()  {
	fmt.Println("2")
}