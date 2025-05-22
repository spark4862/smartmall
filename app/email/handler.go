package main

import (
	"context"
	email "github.com/spark4862/smartmall/rpc_gen/kitex_gen/email"
	"github.com/spark4862/smartmall/app/email/biz/service"
)

// EmailServiceImpl implements the last service interface defined in the IDL.
type EmailServiceImpl struct{}

// SendEmail implements the EmailServiceImpl interface.
func (s *EmailServiceImpl) SendEmail(ctx context.Context, req *email.EmailReq) (resp *email.EmailResp, err error) {
	resp, err = service.NewSendEmailService(ctx).Run(req)

	return resp, err
}
