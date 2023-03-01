package handler

import (
	"net/http"

	"code-push-server/bundle/internal/logic"
	"code-push-server/bundle/internal/svc"
	"code-push-server/bundle/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BundleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBundleLogic(r.Context(), svcCtx)
		resp, err := l.Bundle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
