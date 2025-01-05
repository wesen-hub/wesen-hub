package main

import (
	"awesomeProject/src/api"
	"awesomeProject/src/service/dispatcher"
	"fmt"
)

type person struct {
	Name string
}

func main() {
	server, err := api.Register()
	if server != nil {
		//分发器运行
		dispatcher.Run()
		server.Run(":8081")
	} else {
		fmt.Println("fail to start server, for reason:/n" + err.Error())
	}
}
