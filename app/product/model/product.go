package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name string `json:"name"`
	// 注意:后面不能有空格
	// 导出字段默认序列化且名称不变 omitempty当字段为空时忽略，
	// 导出字段默认映射为gorm列，名称转为snake格式
	// -忽略导出字段，均有效
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Price       int    `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_categories;"`
}

func (p *Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}
