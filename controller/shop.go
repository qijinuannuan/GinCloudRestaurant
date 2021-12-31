package controller

import (
	"gincloudrestaurant/service"
	"gincloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

type ShopController struct {

}

func (sc *ShopController) Router(engine *gin.Engine) {
	engine.GET("/api/shops", sc.ShopList)
	engine.GET("/api/search_shops", sc.SearchShop)
}

func (sc *ShopController) SearchShop(ctx *gin.Context) {
	longitude := ctx.Query("longitude")
	latitude := ctx.Query("latitude")
	keyword := ctx.Query("keyword")
	if keyword == "" {
		tool.Failed(ctx, "input keyword")
		return
	}
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined"{
		longitude = "116.34"
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shopService.SearchShop(longitude, latitude, keyword)
}

func (sc *ShopController) ShopList(ctx *gin.Context) {
	longitude := ctx.Query("longitude")
	latitude := ctx.Query("latitude")
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined"{
		longitude = "116.34"
		latitude = "40.34"
	}
	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) == 0 {
		tool.Failed(ctx, "get nothing of shops")

		return
	}
	for _, shop := range shops {
		shopServices := shopService.GetService(shop.Id)
		if len(shopServices) == 0 {
			shop.Supports = nil
		} else {
			shop.Supports = shopServices
		}
	}
	tool.Success(ctx, shops)
	return
}
