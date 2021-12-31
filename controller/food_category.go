package controller

import (
	"gincloudrestaurant/service"
	"gincloudrestaurant/tool"
	"github.com/gin-gonic/gin"
)

type FoodCategoryController struct {

}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	engine.GET("/api/food_category", fcc.FoodCategory)
}

func (fcc *FoodCategoryController) FoodCategory(ctx *gin.Context) {
	foodCategoryService := service.FoodCategoryService{}
	categories, err := foodCategoryService.Categories()
	if err != nil {
		tool.Failed(ctx, "food category get failed")
		return
	}
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = "图片服务器地址" + "/" + category.ImageUrl
		}
	}
	tool.Success(ctx, categories)
}