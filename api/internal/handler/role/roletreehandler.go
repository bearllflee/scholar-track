package role

import (
	"github.com/bearllflee/scholar-track/pkg/response"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/role"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RoleTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleTreeReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := role.NewRoleTreeLogic(r.Context(), svcCtx)
		resp, err := l.RoleTree(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "获取角色树失败")
		} else {
			response.SuccessWithData(r.Context(), w, resp)
		}
	}
}
