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

func TestGetCart_Run(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	projectRoot := filepath.Join(dir, "..", "..")
	_ = os.Chdir(projectRoot)
	_ = godotenv.Load()
	mysql.Init()
	// 注意此处
	rpc.Init()

	ctx := context.Background()
	s := NewGetCartService(ctx)
	// init req and assert value

	// 不想解耦了，先跑一边add_item_test
	testCases := []uint32{
		1,    // should return 1 item
		1000, // should return none
		0,    // should return none
	}

	for _, tc := range testCases {
		req := &cart.GetCartReq{
			UserId: tc,
		}
		resp, err := s.Run(req)
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}

	// todo: edit your unit test
}
