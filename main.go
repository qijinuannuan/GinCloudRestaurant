package main

import (
	"fmt"
	"gincloudrestaurant/controller"
	"gincloudrestaurant/tool"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main() {
	cfg, err := tool.ParseConfig("./conf/app.json")
	if err != nil {
		panic(err.Error())
	}
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//初始化redis配置
	tool.InitRedisStore()

	app := gin.Default()

	//设置全局跨域访问
	app.Use(Cors())
	tool.InitSession(app)

	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
	new(controller.FoodCategoryController).Router(router)
	new(controller.ShopController).Router(router)
	new(controller.GoodController).Router(router)
}

// Cors 跨域访问：cross origin resource share
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range ctx.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			ctx.Header("Access-Control-Max-Age", "172800")
			ctx.Header("Access-Control-Allow-Credentials", "false")
			ctx.Set("content-type", "application/json") //// 设置返回格式是json
		}
		if method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "Options Request!")
		}
		//继续处理其他请求
		ctx.Next()
	}
}