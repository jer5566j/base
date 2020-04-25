package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var conf *Conf
var confEnv string

func init() {
	flag.StringVar(&confEnv, "C", "dev", "config env, 'dev' or 'test' or 'prod', default 'dev'")
	flag.Parse()

	//读取配置文件
	confPath := fmt.Sprintf("./config/%s-config.yml", confEnv)
	fmt.Println(fmt.Sprintf("读取配置文件: %s", confPath))
	buf, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println(fmt.Sprintf("读取配置文件错误: %s", err.Error()))
		os.Exit(-1)
	}

	conf = &Conf{}
	err = yaml.Unmarshal(buf, conf)
	if nil != err {
		fmt.Println(fmt.Sprintf("解析配置文件错误: %s", err.Error()))
		os.Exit(-1)
	}

	//conf.App.RootPath, _ = filepath.Abs(filepath.Dir(os.Args[0]))
}

func GetConf() Conf {
	return *conf
}
