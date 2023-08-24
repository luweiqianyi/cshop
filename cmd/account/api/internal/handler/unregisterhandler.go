package handler

import (
	"net/http"

	"cshop/cmd/account/api/internal/logic"
	"cshop/cmd/account/api/internal/svc"
	"cshop/cmd/account/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UnRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UnRegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUnRegisterLogic(r.Context(), svcCtx)
		resp, err := l.UnRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
