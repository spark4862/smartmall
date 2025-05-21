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

func TestSearchProducts_Run(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	projectRoot := filepath.Join(dir, "..", "..")
	_ = os.Chdir(projectRoot)
	_ = godotenv.Load()
	mysql.Init()

	ctx := context.Background()
	s := NewSearchProductsService(ctx)
	// init req and assert value

	testCases := []string{
		"egg",
		"rich",
		"agf coffee is tasty", // return none
	}

	for _, tc := range testCases {
		req := &product.SearchProductsReq{
			Query: tc,
		}
		resp, err := s.Run(req)
		t.Logf("err: %v", err)
		t.Logf("resp: %v", resp)
	}

	// todo: edit your unit test
}
