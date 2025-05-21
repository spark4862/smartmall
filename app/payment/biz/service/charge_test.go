package service

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/spark4862/smartmall/app/payment/biz/dal/mysql"
	payment "github.com/spark4862/smartmall/rpc_gen/kitex_gen/payment"
)

func TestCharge_Run(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	// 假设项目根目录在当前文件的上两级目录
	projectRoot := filepath.Join(dir, "..", "..")
	_ = os.Chdir(projectRoot)
	_ = godotenv.Load()
	mysql.Init()

	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{
		Card: &payment.CreditCard{
			Id: "4111111111111111",
			// alter this
			Cvv:             123,
			ExpirationMonth: 1,
			ExpirationYear:  2027,
			// alter this
		},
		Amount:  100,
		OrderId: "123456",
		UserId:  1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
