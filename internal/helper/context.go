package helper

import (
	"context"
	"hermes/internal/constant/types"
	contextmodel "hermes/internal/model/context"
)

func GetUserFromContext(ctx context.Context) (contextmodel.User, bool) {
	user, ok := ctx.Value(types.ContextKeyUser).(contextmodel.User)
	return user, ok
}
