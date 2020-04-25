package api

import (
	"base/config"
	"base/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var router *gin.Engine
var conf = config.GetConf()

func init() {
	gin.DefaultWriter = os.Stdin
	if conf.App.Debug == false {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()

	// 开发环境启用跨域中间件
	if conf.App.Debug {
		router.Use(cors())
	}

	// 日志初始化
	initLog()

	// 路由初始化
	initRoute()
}

func Start() {
	addr := fmt.Sprintf("%s:%s", conf.Http.Ip, conf.Http.Port)
	utils.Log("info", fmt.Sprintf("启动Http服务: %s", addr))
	err := router.Run(addr)
	if nil != err {
		utils.Log("error", fmt.Sprintf("启动HTTP服务失败 %s", err.Error()))
		os.Exit(-1)
	}
}

func initLog() {
	path := fmt.Sprintf("%v/logs/api/", conf.App.RootPath)
	if _, err := os.Stat(path); err != nil {
		err = os.MkdirAll(path, 0666)
		if nil != err {
			utils.Log("error", fmt.Sprintf("初始化HTTP日志目录失败 %s", err.Error()))
			os.Exit(1)
		}
	}

	filePath := fmt.Sprintf("%v%v", path, "error.log")
	if _, err := os.Stat(filePath); err != nil {
		file, err := os.Create(filePath)
		if err != nil {
			utils.Log("error", fmt.Sprintf("初始化HTTP日志文件失败 %s", err.Error()))
			os.Exit(1)
		}
		err = file.Close()
		if nil != err {

		}
	}

	writeFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	router.Use(gin.RecoveryWithWriter(writeFile))
}
