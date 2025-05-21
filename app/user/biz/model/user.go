package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex; type:varchar(255)"`
	// 和unique的区别在于可以创建索引做更精细的控制，如两个字段使用相同带uniqueIndex名，作为唯一字段组
	PasswordHashed string `gorm:"type:varchar(255) not null"`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return &user, err
}
