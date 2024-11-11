package response

import (
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(ctx context.Context, w http.ResponseWriter, code int, data interface{}, message string) {
	resp := &Response{
		Code:    code,
		Data:    data,
		Message: message,
	}
	httpx.OkJsonCtx(ctx, w, resp)
}

func Error(ctx context.Context, w http.ResponseWriter) {
	Result(ctx, w, ERROR, nil, "操作失败")
}

func ErrWithMessage(ctx context.Context, w http.ResponseWriter, message string) {
	Result(ctx, w, ERROR, nil, message)
}

func Success(ctx context.Context, w http.ResponseWriter) {
	Result(ctx, w, SUCCESS, nil, "操作成功")
}

func SuccessWithData(ctx context.Context, w http.ResponseWriter, data interface{}) {
	Result(ctx, w, SUCCESS, data, "查询成功")
}

func SuccessWithDetail(ctx context.Context, w http.ResponseWriter, data interface{}, message string) {
	Result(ctx, w, SUCCESS, data, message)
}

func ErrWithNoAuth(ctx context.Context, w http.ResponseWriter, message string) {
	Result(ctx, w, ERROR, nil, message)
}
