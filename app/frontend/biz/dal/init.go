package dal

import (
	"github.com/spark4862/smartmall/app/frontend/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
