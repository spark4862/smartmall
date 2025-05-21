package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId    uint32 `gorm:"type:int(11) not null; index:idx_user_id"`
	ProductId int32  `gorm:"type:int(11) not null"`
	Qty       int32  `gorm:"type:int(11) not null"`
}

func TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, pItem *Cart) error {
	var item Cart
	err := db.WithContext(ctx).Model(&Cart{}).
		Where(&Cart{UserId: pItem.UserId, ProductId: pItem.ProductId}).
		First(&item).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if item.ID > 0 {
		return db.WithContext(ctx).Model(&Cart{}).
			Where(&Cart{UserId: pItem.UserId, ProductId: pItem.ProductId}).
			Update("qty", gorm.Expr("qty+?", pItem.Qty)).Error
		// 这里不直接update可以防止并发问题
		// 如果用？的占位符方式，比如gorm.Expr,update,预编译都只会编译一条，对于不同qty和
		// UPDATE cart SET qty = %d WHERE user_id = %d AND product_id = %d 这种写死方式则会编译多条
	}

	return db.Create(pItem).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("userId is empty")
	}
	return db.Where("user_id = ?", userId).Delete(&Cart{}, "user_id = ?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	if userId == 0 {
		return nil, errors.New("userId is empty")
	}
	var items []*Cart
	err := db.WithContext(ctx).Model(&Cart{}).
		Where(&Cart{UserId: userId}).Find(&items).Error
	return items, err
}
