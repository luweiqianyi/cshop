package handler

import (
	"net/http"

	"cshop/cmd/order/api/internal/logic"
	"cshop/cmd/order/api/internal/svc"
	"cshop/cmd/order/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CommitOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommitOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCommitOrderLogic(r.Context(), svcCtx)
		resp, err := l.CommitOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
