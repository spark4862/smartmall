package service

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
	"github.com/spark4862/smartmall/app/product/biz/dal/mysql"
	product "github.com/spark4862/smartmall/rpc_gen/kitex_gen/product"
)

func TestGetProduct_Run(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	projectRoot := filepath.Join(dir, "..", "..")
	_ = os.Chdir(projectRoot)
	_ = godotenv.Load()
	mysql.Init()

	ctx := context.Background()
	s := NewGetProductService(ctx)
	// init req and assert value

	testCases := []int32{
		1,    // should return product with id 1
		1000, // should return err
		0,    // should return err
	}

	for _, tc := range testCases {
		req := &product.GetProductReq{
			Id: tc,
		}
		resp, err := s.Run(req)
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}

	// todo: edit your unit test
}
