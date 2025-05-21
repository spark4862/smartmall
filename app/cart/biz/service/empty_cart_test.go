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

func TestEmptyCart_Run(t *testing.T) {
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

	req := &cart.AddItemReq{
		UserId: 1,
		Item: &cart.CartItem{
			ProductId: 1,
			Quantity:  1,
		},
	}
	_, _ = s.Run(req)

	testCases := []uint32{
		1,    // 正常值
		1000, // 异常
		0,    // 异常
		1,    // 重复删除
	}

	s1 := NewEmptyCartService(ctx)

	for _, tc := range testCases {
		req := &cart.EmptyCartReq{
			UserId: tc,
		}
		resp, err := s1.Run(req)
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}

	// todo: edit your unit test
}
