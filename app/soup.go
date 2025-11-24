package app

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type soupHandler struct {
}

// 首页
var router = make([]string, 0, 10000)

func (handler *soupHandler) index(context *gin.Context) {
	resp := response.GenericResponse(http.StatusOK, "SUCCESS", router[rand.Intn(len(router))])
	context.HTML(http.StatusOK, "index.html", resp)
}

func init() {
	rand.Seed(time.Now().UnixNano())
	fileReader, err := os.Open("conf/soup.txt")
	if err != nil {
		panic(err)
		return
	}

	scanner := bufio.NewScanner(fileReader)
	scannerErr := scanner.Err()
	if scannerErr != nil {
		panic(err)
		return
	}

	for scanner.Scan() {
		router = append(router, scanner.Text())
	}
}

// InitSoupHandler 加载路由
func InitSoupHandler() {
	// 组件初始化
	handler := soupHandler{}

	// 路由初始化
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "", HandlerFunc: handler.index})
	server.RegisterRoute("/", route)
}
