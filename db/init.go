package db

import (
	"base/config"
	"base/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	"xorm.io/core"
	"xorm.io/xorm"
)

var conf = config.GetConf()
var engine *xorm.Engine

func init() {
	utils.Log("info", "初始化数据库...")
	var err error

	// 创建数据库引擎
	engine, err = xorm.NewEngine(conf.DB.Type, conf.DB.Dsn)
	if nil != err {
		utils.Log("error", fmt.Sprintf("数据库初始化失败: %s", err.Error()))
		os.Exit(1)
	}

	// 测试连接
	if err := engine.Ping(); nil != err {
		utils.Log("error", fmt.Sprintf("测试数据库链接失败: %s", err.Error()))
		os.Exit(1)
	}

	engine.ShowSQL(conf.App.Debug)
	// 设置最大链接数
	engine.SetMaxOpenConns(conf.DB.MaxOpenConnect)
	// 设置最大闲置连接数
	engine.SetMaxIdleConns(conf.DB.MaxIdleConnect)
	// 设置链接最大时间
	engine.SetConnMaxLifetime(time.Second * 12 * 60 * 60)
	// 设置表前缀
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, conf.DB.TablePrefix)
	engine.SetTableMapper(tbMapper)
}

// 返回xorm实例
func GetDB() *xorm.Engine {
	return engine
}
