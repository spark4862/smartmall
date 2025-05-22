package service

import (
	"context"
	email "github.com/spark4862/smartmall/rpc_gen/kitex_gen/email"
)

type SendEmailService struct {
	ctx context.Context
} // NewSendEmailService new SendEmailService
func NewSendEmailService(ctx context.Context) *SendEmailService {
	return &SendEmailService{ctx: ctx}
}

// Run create note info
func (s *SendEmailService) Run(req *email.EmailReq) (resp *email.EmailResp, err error) {
	// Finish your business logic.

	return
}
