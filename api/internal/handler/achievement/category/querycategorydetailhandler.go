package category

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/achievement/category"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryCategoryDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryCategoryDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewQueryCategoryDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryCategoryDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
