package service

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/spark4862/smartmall/app/cart/biz/dal/mysql"
	"github.com/spark4862/smartmall/app/cart/rpc"
	cart "github.com/spark4862/smartmall/rpc_gen/kitex_gen/cart"
)

func TestAddItem_Run(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	projectRoot := filepath.Join(dir, "..", "..")
	_ = os.Chdir(projectRoot)
	_ = godotenv.Load()
	mysql.Init()
	// 注意此处
	rpc.Init()

	ctx := context.Background()
	s := NewAddItemService(ctx)
	// init req and assert value

	testCases := []int32{
		1,    // should return none
		1000, // should return err
		0,    // should return err
		1,    // check db , quantity should be 2
	}

	for _, tc := range testCases {
		req := &cart.AddItemReq{
			UserId: 1,
			Item: &cart.CartItem{
				ProductId: tc,
				Quantity:  1,
			},
		}
		resp, err := s.Run(req)
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}
	// todo: edit your unit test
}
