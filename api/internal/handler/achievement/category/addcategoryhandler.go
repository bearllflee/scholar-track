package category

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/achievement/category"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddCategoryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCategoryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := category.NewAddCategoryLogic(r.Context(), svcCtx)
		resp, err := l.AddCategory(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
