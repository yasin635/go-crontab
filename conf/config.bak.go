package conf

import (
	"flag"
	"github.com/larspensjo/config"
	"runtime"
	"log"
	"fmt"
)

var (
	configFile = flag.String("configfile", "crontab/config.ini", "General configuration file")
)

//topic list
var CONFIG_MYSQL = make(map[string]string)

var CONFIG_REDIS = make(map[string]string)

func main()  {
	fmt.Print("sdsf")
}

func init()  {
	//mysql配置
	mysql_config()
	//redis配置
	//redis_config();
	fmt.Println(configFile)

}
//noinspection GoUnresolvedReference

func mysql_config()  {
	//定义项目跟目录
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("mysql") {
		section, err := cfg.SectionOptions("mysql")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("mysql", v)
				if err == nil {
					CONFIG_MYSQL[v] = options
				}
			}
		}
	}
}
func redis_config()  {
	//定义项目跟目录
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()
	//set config file std
	cfg, err := config.ReadDefault(*configFile)
	if err != nil {
		log.Fatalf("Fail to find", *configFile, err)
	}
	//set config file std End

	//Initialized topic from the configuration
	if cfg.HasSection("redis") {
		section, err := cfg.SectionOptions("redis")
		if err == nil {
			for _, v := range section {
				options, err := cfg.String("redis", v)
				if err == nil {
					CONFIG_REDIS[v] = options
				}
			}
		}
	}
}