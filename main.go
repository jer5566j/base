package main

import (
	"base/api"
	"base/utils"
	"fmt"
	"os"
)

func main()  {
	defer func() {
		if err := recover(); err != nil {
			utils.Log("info", fmt.Sprintf("未捕获的异常: %s", err))
			os.Exit(-1)
		}
	}()

	api.Start()
}
