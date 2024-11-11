package user

import (
	"github.com/bearllflee/scholar-track/api/internal/utils"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/user"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetSelfInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetSelfInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error("参数错误", err)
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}
		req.ID = utils.GetUserId(r)
		l := user.NewSetSelfInfoLogic(r.Context(), svcCtx)
		_, err := l.SetSelfInfo(&req)
		if err != nil {
			logx.Error("更新失败", err)
			response.ErrWithMessage(r.Context(), w, "更新失败")
		} else {
			response.Success(r.Context(), w)
		}
	}
}
