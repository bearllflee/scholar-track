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
	err := IfUniqueHasExists(userModel)
	if err != nil {
		return nil, err
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

func IfUniqueHasExists(user *model.User) error {
	var c int64
	var err error
	if user.ID != 0 {
		err = global.DB.Model(user).Where("username = ? and id != ?", user.Username, user.ID).Count(&c).Error
	} else {
		err = global.DB.Model(user).Where("username = ?", user.Username).Count(&c).Error
	}
	if err != nil {
		return err
	}
	if c > 0 {
		return cerror.ErrUserHasExists
	}
	if user.ID != 0 {
		err = global.DB.Model(user).Where("email = ? and id != ?", user.Email, user.ID).Count(&c).Error
	} else {
		err = global.DB.Model(user).Where("email = ?", user.Email).Count(&c).Error
	}
	if err != nil {
		return err
	}
	if c > 0 {
		return cerror.ErrEmailHasExists
	}
	if user.ID != 0 {
		err = global.DB.Model(user).Where("phone = ? and id != ?", user.Phone, user.ID).Count(&c).Error
	} else {
		err = global.DB.Model(user).Where("phone = ?", user.Phone).Count(&c).Error
	}
	if err != nil {
		return err
	}
	if c > 0 {
		return cerror.ErrPhoneHasExists
	}
	return nil
}
