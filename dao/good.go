package dao

import (
	"gincloudrestaurant/model"
	"gincloudrestaurant/tool"
)

type GoodDao struct {
	*tool.Orm
}

func NewGoodDao() *GoodDao {
	return &GoodDao{Orm:tool.DbEngine}
}

func (gg *GoodDao) QueryGoods(shop_id int64) ([]model.Goods, error) {
	var goods []model.Goods
	err := gg.Orm.Where(" shop_id = ? ", shop_id).Find(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}