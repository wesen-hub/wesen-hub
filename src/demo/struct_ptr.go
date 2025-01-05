package demo

import "github.com/gin-gonic/gin"

// 结构体指语法糖
func main1() {
	// 创建一个默认的Gin引擎
	r := gin.Default()

	// 定义一个路由处理器
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, Gin!"})
	})

	// 启动服务器，监听端口8080
	r.Run(":8080")
}
