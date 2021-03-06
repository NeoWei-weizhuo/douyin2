package main

import (
	"douyin/conf"
	"github.com/gin-gonic/gin"
)

func main() {
	//实例化一个gin对象
	r := gin.Default()

	go initRouter(r)
	conf.Init()
	conf.Redis("127.0.0.1:6379", "")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
