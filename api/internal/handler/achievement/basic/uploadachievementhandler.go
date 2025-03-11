package basic

import (
	"github.com/bearllflee/scholar-track/pkg/response"
	"net/http"

	"github.com/bearllflee/scholar-track/api/internal/logic/achievement/basic"
	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadAchievementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadAchievementReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ErrWithMessage(r.Context(), w, "参数错误")
			return
		}

		materials := r.MultipartForm.File["materials"]
		if len(materials) == 0 {
			response.ErrWithMessage(r.Context(), w, "请上传证明材料")
			return
		}

		for _, material := range materials {
			fileBasic := &types.FileBasic{}
			fileBasic.Name = material.Filename
			fileBasic.Usage = "material"
			fileBasic.Size = material.Size
			fileBytes, err := material.Open()
			if err != nil {
				response.ErrWithMessage(r.Context(), w, "打开文件失败")
				return
			}
			n, err := fileBytes.Read(fileBasic.Content)
			if err != nil || n != int(fileBasic.Size) {
				response.ErrWithMessage(r.Context(), w, "文件读取失败")
				return
			}
			req.Files = append(req.Files, fileBasic)
		}
		pictures := r.MultipartForm.File["pictures"]
		if len(pictures) > 0 {
			for _, picture := range pictures {
				fileBasic := &types.FileBasic{}
				fileBasic.Name = picture.Filename
				fileBasic.Usage = "picture"
				fileBasic.Size = picture.Size
				fileBytes, err := picture.Open()
				if err != nil {
					response.ErrWithMessage(r.Context(), w, "打开文件失败")
					return
				}
				n, err := fileBytes.Read(fileBasic.Content)
				if err != nil || n != int(fileBasic.Size) {
					response.ErrWithMessage(r.Context(), w, "文件读取失败")
					return
				}
				req.Files = append(req.Files, fileBasic)
			}
		}

		l := basic.NewUploadAchievementLogic(r.Context(), svcCtx)
		_, err := l.UploadAchievement(&req)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "上传失败")
		} else {
			response.Success(r.Context(), w)
		}
	}
}
