package biz

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"

	"github.com/rs/xid"
	"kratos-application/api/gen/api"
	"kratos-application/api/gen/v1api"
	"kratos-application/app/data"
	"kratos-application/core/config"
	"kratos-application/core/jwtx"
	"kratos-application/core/utils"
	"kratos-application/util"
)

// AuthRegister 注册用户
func (b *Biz) AuthRegister(ctx context.Context, req *v1api.AuthRegisterRequest) (*v1api.AuthRegisterResponse, error) {
	u, err := b.data.User.FirstEmail(req.Email)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("FirstEmail %v", err)
			return nil, util.ServiceError(err)
		}
	}
	if u.ID > 0 {
		return nil, util.NewError(api.ECode_EEmailExist, nil)
	}

	password, err := utils.BcryptPasswordHash(req.Password)
	if err != nil {
		log.Errorf("FirstEmail %v", err)
		return nil, util.ServiceError(err)
	}
	user := &data.User{
		Uname:     xid.New().String(),
		Email:     req.Email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = b.data.User.Create(user)
	if err != nil {
		log.Errorf("Create %v", err)
		return nil, util.ServiceError(err)
	}

	token, err := jwtx.GenToken(user.ID, config.Get().Jwt)
	if err != nil {
		log.Errorf("GenToken %v", err)
		return nil, util.ServiceError(err)
	}
	resp := &v1api.AuthRegisterResponse{
		Token: token,
		User: &v1api.User{
			Uid:   user.ID,
			Uname: user.Uname,
			Email: user.Email,
		},
	}
	return resp, nil
}

// AuthLogin 登录用户
func (b *Biz) AuthLogin(ctx context.Context, req *v1api.AuthLoginRequest) (*v1api.AuthLoginResponse, error) {
	user, err := b.data.User.FirstEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util.NewError(api.ECode_EEmailNotExist, nil)
		}
		log.Errorf("FirstEmail %v", err)
		return nil, util.ServiceError(err)
	}
	if !utils.BcryptPasswordCheck(req.Password, user.Password) {
		return nil, util.NewError(api.ECode_EPasswordError, nil)
	}
	token, err := jwtx.GenToken(user.ID, config.Get().Jwt)
	if err != nil {
		log.Errorf("GenToken %v", err)
		return nil, util.ServiceError(err)
	}
	resp := &v1api.AuthLoginResponse{
		Token: token,
		User: &v1api.User{
			Uid:   user.ID,
			Uname: user.Uname,
			Email: user.Email,
		},
	}
	return resp, nil
}

// AuthRefresh 刷新Jwt
func (b *Biz) AuthRefresh(ctx context.Context, req *v1api.AuthRefreshRequest) (*v1api.AuthRefreshResponse, error) {
	uid, ok := jwtx.FromContextUid(ctx)
	if !ok {
		return nil, util.NewError(api.ECode_EAuthExpire, nil)
	}
	token, err := jwtx.GenToken(uid, config.Get().Jwt)
	if err != nil {
		log.Errorf("GenToken %v", err)
		return nil, util.ServiceError(err)
	}
	resp := &v1api.AuthRefreshResponse{
		Token: token,
	}
	return resp, nil
}
