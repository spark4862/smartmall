package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/spark4862/smartmall/app/frontend/hertz_gen/frontend/auth"
	"github.com/spark4862/smartmall/app/frontend/infra/rpc"
	"github.com/spark4862/smartmall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	if err = session.Save(); err != nil {
		return
	}
	if req.Next != "" {
		redirect = req.Next
	} else {
		redirect = "/"
	}
	return
}
