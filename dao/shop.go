package dao

import (
	"gincloudrestaurant/model"
	"gincloudrestaurant/tool"
)

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{Orm: tool.DbEngine}
}

const DEFAULT_RANGE = 5

func (sd *ShopDao) QueryShops(longitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop
	if keyword == "" {
		err := sd.Engine.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ? and status = 1",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
		if err != nil {
			return nil
		}
	} else {
		err := sd.Engine.Where(" longitude > ? and longitude < ? and latitude > ? and latitude < ? and name like ? and status = 1",
			longitude-DEFAULT_RANGE, longitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, keyword).Find(&shops)
		if err != nil {
			return nil
		}
	}
	return shops
}

func (sd *ShopDao) QueryServiceByShopId(shopId int64) []model.Service {
	var services []model.Service
	err := sd.Orm.Table("service").Join("INNER", "shop_service", " service.id = shop_service.service_id and shop_service.shop_id = ?", shopId).Find(&services)
	if err != nil {
		return nil
	}
	return services
}
