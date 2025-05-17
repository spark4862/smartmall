package utils

import "context"

func GetUserIdFromCtx(ctx context.Context) int32 {
	userId, ok := ctx.Value(SessionUserId).(int32)
	if !ok {
		return 0
	}
	return userId
}
