package service

import (
	"gincloudrestaurant/dao"
	"gincloudrestaurant/model"
)

type GoodService struct {

}

func NewGoodService() *GoodService {
	return &GoodService{}
}

func (gs *GoodService) QueryGoods(shop_id int64) []model.Goods {
	goodDao := dao.NewGoodDao()
	goods, err := goodDao.QueryGoods(shop_id)
	if err != nil {
		return nil
	}
	return goods
}