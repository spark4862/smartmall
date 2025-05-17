package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

// context包中不推荐使用string类型作为key,以减少碰撞风险

// 这里需要注意，由于map中key是interface{}，比较时会比较类型和key

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, reqCtx *app.RequestContext) {
		s := sessions.Default(reqCtx)
		ctx = context.WithValue(ctx, SessionUserId, s.Get("user_id"))

		reqCtx.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, reqCtx *app.RequestContext) {
		s := sessions.Default(reqCtx)
		userId := s.Get("user_id")
		if userId == nil {
			reqCtx.Redirect(consts.StatusFound, []byte("/sign-in?next="+reqCtx.FullPath()))
			// abort 后续handler
			reqCtx.Abort()
			return
		}
		reqCtx.Next(ctx)
	}
}
