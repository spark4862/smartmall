package service

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/spark4862/smartmall/app/checkout/infra/rpc"
	checkout "github.com/spark4862/smartmall/rpc_gen/kitex_gen/checkout"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/payment"
)

func TestCheckout_Run(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	// 假设项目根目录在当前文件的上两级目录
	projectRoot := filepath.Join(dir, "..", "..")
	_ = os.Chdir(projectRoot)
	// _ = godotenv.Load()
	rpc.Init()

	ctx := context.Background()
	s := NewCheckoutService(ctx)
	// init req and assert value

	req := &checkout.CheckoutReq{
		UserId: 1,
		Address: &checkout.Address{
			Street:  "street",
			City:    "city",
			State:   "state",
			Country: "country",
			Zip:     "zip",
		},
		CreditCard: &payment.CreditCard{
			Id: "4111111111111111",
			// alter this
			Cvv:             123,
			ExpirationMonth: 1,
			ExpirationYear:  2027,
			// alter this
		},
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email@email.com",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
}
