package server

import (
	"github.com/gin-gonic/gin"
)

type Server = *gin.Engine //这里注意区分和type server *gin.Engine 的区别

func Init() Server {
	var instance Server

	// 创建一个默认的Gin引擎
	instance = gin.Default()
	//todo 这里添加初始化逻辑

	return instance
}
