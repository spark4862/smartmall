package service

import (
	"context"
	"testing"
	order "github.com/spark4862/smartmall/rpc_gen/kitex_gen/order"
)

func TestListOrders_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOrdersService(ctx)
	// init req and assert value

	req := &order.ListOrdersReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
