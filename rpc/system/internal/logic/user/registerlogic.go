package userlogic

import (
	"context"
	"github.com/bearllflee/scholar-track/pkg/cerror"
	"github.com/bearllflee/scholar-track/pkg/global"
	"github.com/bearllflee/scholar-track/rpc/system/client/user"
	"github.com/bearllflee/scholar-track/rpc/system/internal/model"
	"github.com/bearllflee/scholar-track/rpc/system/internal/svc"
	"github.com/bearllflee/scholar-track/rpc/system/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *system.RegisterReq) (*system.RegisterResp, error) {
	var c int64
	temp := &model.User{}
	err := global.DB.Model(temp).Where("username = ?", in.Username).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, cerror.ErrUserHasExists
	}
	err = global.DB.Model(temp).Where("email = ?", in.Email).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, cerror.ErrEmailHasExists
	}
	err = global.DB.Model(temp).Where("phone = ?", in.Phone).Count(&c).Error
	if err != nil {
		return nil, err
	}
	if c > 0 {
		return nil, cerror.ErrPhoneHasExists
	}
	userModel := &model.User{
		Username: in.Username,
		Email:    in.Email,
		Role:     uint(in.Role),
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Gender:   int8(in.Gender),
		Major:    in.Major,
		College:  in.College,
		Grade:    in.Grade,
		Class:    in.Class,
		Realname: in.Realname,
		Password: in.Password,
	}
	err = global.DB.Create(userModel).Error
	if err != nil {
		return nil, err
	}
	return &system.RegisterResp{
		User: &user.QueryUserDetailResp{
			Id:        int64(userModel.ID),
			CreatedAt: userModel.CreatedAt.Unix(),
			UpdatedAt: userModel.UpdatedAt.Unix(),
			Username:  userModel.Username,
			Email:     userModel.Email,
			Role:      int64(userModel.Role),
			Nickname:  userModel.Nickname,
			Phone:     userModel.Phone,
			Gender:    int32(userModel.Gender),
			Major:     userModel.Major,
			College:   userModel.College,
			Grade:     userModel.Grade,
			Class:     userModel.Class,
			Realname:  userModel.Realname,
		},
	}, nil
}
