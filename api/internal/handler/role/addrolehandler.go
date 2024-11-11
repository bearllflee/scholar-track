package role

import (
	"github.com/bearllflee/scholar-track/pkg/response"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/role"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := role.NewAddRoleLogic(r.Context(), svcCtx)
		resp, err := l.AddRole(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "添加角色失败")
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
