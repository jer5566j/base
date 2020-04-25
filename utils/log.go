package utils

import (
	"base/config"
	"fmt"
	"os"
	"runtime"
	"time"
)

// 写入日志
func Log(level string, msg string, rate ...int) {
	skip := 2
	if nil != rate {
		skip = rate[0]
	}

	func(lv string, mg string, sk int) {
		// 非debug模式不记录debug级别日志
		if config.GetConf().App.Debug == false && lv == "debug" {
			return
		}

		nowTime := time.Now()
		dateStr := nowTime.Format("20060102")
		timeStr := nowTime.Format("15")
		rootPath := config.GetConf().App.RootPath
		logDir := fmt.Sprintf("%s/logs/%s/", rootPath, dateStr)
		if _, er := os.Stat(logDir); er != nil {
			os.MkdirAll(logDir, 0666)
		}

		file := logDir + timeStr + ".log"
		if _, er := os.Stat(file); er != nil {
			f, _ := os.Create(file)
			f.Close()
		}

		writeFile, er := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if er != nil {
			fmt.Println(er)
		}

		timeLog := nowTime.Format("2006-01-02 15:04:05")
		funcName, _, line, _ := runtime.Caller(sk)
		format := "%v [%s: %v] [%s] ===>> %s\r\n"
		logStr := fmt.Sprintf(format, timeLog, runtime.FuncForPC(funcName).Name(), line, lv, mg)

		fmt.Print(logStr)
		_, ers := writeFile.WriteString(logStr)
		if ers != nil {
			fmt.Println(ers)
		}

		defer writeFile.Close()
	}(level, msg, skip)
}

func SqlError(msg string, sqlStr string, param interface{}) {
	logStr := fmt.Sprintf("%s \n[SQL] %s\n[PARAM] %v", msg, sqlStr, param)
	Log("error", logStr, 3)
}
