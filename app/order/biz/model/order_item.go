package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderId string `gorm:"type:varchar(100); uniqueIndex"`
	// 定义：INT(M) 中的 M 表示显示宽度，用于指定在某些情况下（如使用 ZEROFILL 属性时）数值的最小显示宽度。
	// 需要注意的是，显示宽度并不影响实际存储的数值范围。对于 INT 类型，无论显示宽度如何设置，其存储范围始终是固定的
	// 显示宽度的废弃：从 MySQL 8.0.17 版本开始，整数类型的显示宽度（如 INT(11)）已被废弃，定义为 INT 即可。
	// 总结：别学别用
	ProductId int32   `gorm:"type:int(11)"`
	Quantity  int32   `gorm:"type:int(11)"`
	Cost      float32 `gorm:"type:decimal(10,2)"`
	//DECIMAL 类型提供精确的数值表示，避免了浮点数可能出现的精度误差
	//：与 FLOAT 或 DOUBLE 相比，DECIMAL 在某些情况下可能会有较低的性能，但提供了更高的精度
	// 位数（M）：10 表示该数值最多可以包含 10 位数字（不包括小数点和符号位）
}

func (OrderItem) TableName() string {
	return "order_item"
}
