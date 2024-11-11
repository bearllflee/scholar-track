package user

import (
	"errors"
	"github.com/bearllflee/scholar-track/pkg/response"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/user"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryUserDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryUserDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		l := user.NewQueryUserDetailLogic(r.Context(), svcCtx)
		resp, err := l.QueryUserDetail(&req)
		if err != nil {
			if errors.Is(err, sqlc.ErrNotFound) {
				response.ErrWithMessage(r.Context(), w, "用户不存在")
			} else {
				response.ErrWithMessage(r.Context(), w, "查询用户失败")
			}
		} else {
			response.SuccessWithDetail(r.Context(), w, resp, "查询用户成功")
		}
	}
}
