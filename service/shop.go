package service

import (
	"gincloudrestaurant/dao"
	"gincloudrestaurant/model"
	"strconv"
)

type ShopService struct {

}

func (ss *ShopService) SearchShop(long, lat, keyword string) []model.Shop {
	longtitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}
	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longtitude, latitude, keyword)
}

func (ss *ShopService) ShopList(long, lat string) []model.Shop {
	longtitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		return nil
	}
	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longtitude, latitude, "")
}

func (ss *ShopService) GetService(shopId int64) []model.Service {
	shopDao := dao.NewShopDao()
	return shopDao.QueryServiceByShopId(shopId)
}