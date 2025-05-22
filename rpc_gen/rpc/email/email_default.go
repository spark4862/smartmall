package email

import (
	"context"
	email "github.com/spark4862/smartmall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendEmail(ctx context.Context, req *email.EmailReq, callOptions ...callopt.Option) (resp *email.EmailResp, err error) {
	resp, err = defaultClient.SendEmail(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendEmail call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
