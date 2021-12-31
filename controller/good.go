package controller

import (
	"gincloudrestaurant/service"
	"gincloudrestaurant/tool"
	"github.com/gin-gonic/gin"
	"strconv"
)

type GoodController struct {

}

func (gc *GoodController) Router(engine *gin.Engine) {
	engine.GET("/api/goods", gc.GetGoods)
}

func (gc *GoodController) GetGoods(ctx *gin.Context) {
	shopId, exist := ctx.GetQuery("shop_id")
	if !exist {
		tool.Failed(ctx, "param parse failed")
		return
	}
	id, err := strconv.Atoi(shopId)
	if err != nil {
		tool.Failed(ctx, "param error")
		return
	}
	goodService := service.NewGoodService()
	goods := goodService.QueryGoods(int64(id))
	if len(goods) == 0 {
		tool.Failed(ctx, "find nothing")
		return
	}
	tool.Success(ctx, goods)
	return
}