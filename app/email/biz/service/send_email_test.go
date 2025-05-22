package service

import (
	"context"
	"testing"
	email "github.com/spark4862/smartmall/rpc_gen/kitex_gen/email"
)

func TestSendEmail_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendEmailService(ctx)
	// init req and assert value

	req := &email.EmailReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
