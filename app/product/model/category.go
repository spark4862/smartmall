package model

type Category struct {
	Base
	Name        string `json:"name"`
	Description string `json:"description"`

	Products []Product `json:"products" gorm:"many2many:product_categories;"`
}

func (c *Category) TableName() string {
	return "category"
}
