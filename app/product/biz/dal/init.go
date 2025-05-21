package dal

import (
	"github.com/spark4862/smartmall/app/product/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
