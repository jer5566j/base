package api

func initRoute() {
	// 静态文件地址
	router.Static("/static", "./static")
}
