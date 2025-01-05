package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "awesomeProject/src/server"

type api struct {
	path       string
	handler    gin.HandlerFunc
	httpMethod string
}

func apis() []api {
	return []api{
		{"/", sayHello, http.MethodGet},
		{"/text/v1", sentenceToWord, http.MethodPost},
		{"/label/v2", textToSentence, http.MethodPost},
	}
}
func instances() *gin.Engine {
	return server.Init()
}

func Register() (server.Server, error) {
	server := instances()
	for _, api := range apis() {
		switch api.httpMethod {
		case http.MethodGet:
			server.GET(api.path, api.handler)
		case http.MethodPost:
			server.POST(api.path, api.handler)
		case http.MethodDelete:
			server.DELETE(api.path, api.handler)
		case http.MethodHead:
			server.HEAD(api.path, api.handler)
		case http.MethodPut:
			server.PUT(api.path, api.handler)
		case http.MethodOptions:
			server.OPTIONS(api.path, api.handler)
		case http.MethodPatch:
			server.PATCH(api.path, api.handler)
		default:
			fmt.Println("请求方法不支持")
			return nil, errors.New("unsupported method")
		}
	}
	return server, nil
}

func sentenceToWord(context *gin.Context) {
	//todo 业务逻辑
}
func textToSentence(context *gin.Context) {

}
func sayHello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello world")
}
