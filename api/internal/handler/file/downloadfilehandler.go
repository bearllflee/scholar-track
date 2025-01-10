package file

import (
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/file"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DownloadFileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DownloadFileReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(r.Context(), "解析请求参数失败", err)
			response.ErrWithMessage(r.Context(), w, "解析请求参数失败")
			return
		}

		l := file.NewDownloadFileLogic(r.Context(), svcCtx)
		resp, err := l.DownloadFile(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, err.Error())
		} else {
			response.SuccessWithFile(r.Context(), w, resp.FileData, resp.FileName, resp.FileType, resp.FileSize)
		}
	}
}

