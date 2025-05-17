package mysql

import (
	"fmt"
	"os"

	"github.com/spark4862/smartmall/app/user/biz/model"
	"github.com/spark4862/smartmall/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
			// GORM会为每个 SQL 操作创建并缓存预编译语句（prepared statement），以提高数据库操作的性能。
			// 每个不同的 SQL 语句都会被缓存，如果 SQL 语句种类繁多，可能导致内存占用增加
			SkipDefaultTransaction: true,
			// 默认在每次写操作（创建、更新、删除）时，自动开启一个事务，以确保数据的一致性，设置后GORM 将不会自动开启事务。
		},
	)
	_ = DB.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
