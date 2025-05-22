package dal

import (
	"github.com/spark4862/smartmall/app/email/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
