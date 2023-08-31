package middleware

import (
	"context"
	"net/http"
)

const (
	AccountName = "AccountName"
)

type IdentityMiddleware struct {
}

func NewIdentityMiddleware() *IdentityMiddleware {
	return &IdentityMiddleware{}
}

func (m *IdentityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		aName := request.Header.Get(AccountName)
		ctx := request.Context()
		ctx = context.WithValue(ctx, AccountName, aName)
		next(writer, request.WithContext(ctx))
	}
}
