package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	Base
	Name string `json:"name"`
	// 注意:后面不能有空格
	Description string `json:"description"`
	// 结构体标签语法

	Picture string  `json:"picture"`
	Price   float32 `json:"price"`

	Categories []Category `json:"categories" gorm:"many2many:product_category;"`
}

func (p *Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, db *gorm.DB) *ProductQuery {
	return &ProductQuery{
		ctx: ctx,
		db:  db,
	}
}

func (p ProductQuery) GetByID(id int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, id).Error
	return
}

func (p ProductQuery) SearchProducts(query string) (products []*Product, err error) {
	// gorm的Find方法可以接收[]Product和[]*Product作为参数，
	// []*的性能可能更好？原因是[]Product需要吧Product append到结构体后者只需要append指针
	err = p.db.WithContext(p.ctx).Model(&Product{}).
		Where("name LIKE ? or description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&products).Error
	return
}
