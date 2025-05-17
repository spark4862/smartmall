package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/common"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	rpcResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	if err != nil {
		return
	}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", rpcResp.UserId)
	if err = session.Save(); err != nil {
		return
	}

	return
}
