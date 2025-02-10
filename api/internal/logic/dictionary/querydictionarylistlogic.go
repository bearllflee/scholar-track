package dictionary

import (
	"context"

	"github.com/bearllflee/scholar-track/api/internal/svc"
	"github.com/bearllflee/scholar-track/api/internal/types"
	"github.com/bearllflee/scholar-track/rpc/system/client/dictionaryservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryDictionaryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryDictionaryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryDictionaryListLogic {
	return &QueryDictionaryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryDictionaryListLogic) QueryDictionaryList(req *types.QueryDictionaryListReq) (resp *types.QueryDictionaryListResp, err error) {
	rpcResp, err := l.svcCtx.DictionaryService.QueryAllDictionary(l.ctx, &dictionaryservice.QueryAllDictionaryReq{})
	if err != nil {
		return nil, err
	}

	var dictionaries []*types.Dictionary
	for _, dictionary := range rpcResp.Dictionaries {
		dictionaries = append(dictionaries, &types.Dictionary{
			Id:        dictionary.Id,
			Name:      dictionary.Name,
			Type:      dictionary.Type,
			Status:    dictionary.Status,
			Desc:      dictionary.Desc,
			CreatedAt: dictionary.CreatedAt,
			UpdatedAt: dictionary.UpdatedAt,
		})
	}

	return &types.QueryDictionaryListResp{
	}, nil
}
