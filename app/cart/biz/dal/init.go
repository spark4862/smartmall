package dal

import (
	"github.com/spark4862/smartmall/app/cart/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
