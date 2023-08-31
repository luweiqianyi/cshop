package ctxdata

import "context"

const (
	AccountName = "accountName"
)

func GetAccountNameFromCtx(ctx context.Context) string {
	accountName, ok := ctx.Value(AccountName).(string)
	if !ok {
		return ""
	}
	return accountName
}
