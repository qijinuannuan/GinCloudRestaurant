package model

type FoodCategory struct {
	Id          int64  `xorm:"pk autoincr" json:"id"`
	Title       string `xorm:"varchar(20)" json:"title"`
	Description string `xorm:"varchar(30)" json:"description"`
	ImageUrl    string `xorm:"varchar(255)" json:"image_url"`
	LinkUrl     string `xorm:"varchar(255)" json:"link_url"`
	IsInServing bool   `json:"is_in_serving"`
}
