package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email   string
	Street  string
	City    string
	State   string
	Country string
	ZipCode string
}

type Order struct {
	gorm.Model
	OrderId    string      `gorm:"type:varchar(100); uniqueIndex"`
	UserId     uint32      `gorm:"type:int(11)"`
	Currency   string      `gorm:"type:varchar(10)"`
	Consignee  Consignee   `gorm:"embedded"` // 用于嵌入结构体
	OrderItems []OrderItem `gorm:"foreignKey:OrderId; references:OrderId;"`
	// https://gorm.io/zh_CN/docs/has_many.html
}

func (Order) TableName() string {
	return "order"
}

func ListOrder(ctx context.Context, db *gorm.DB, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.WithContext(ctx).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
