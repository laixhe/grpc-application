package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"kratos-application/api/gen/auth"
	"kratos-application/app/cache"
	"kratos-application/app/data"
	"kratos-application/core/errorx"
	"kratos-application/core/i18nx"
	"kratos-application/core/jwtx"
)

type AuthBiz struct {
	Data  *data.Data
	Cache *cache.Cache
}

func NewAuthBiz() *AuthBiz {
	return &AuthBiz{
		Data:  newBiz.Data,
		Cache: newBiz.Cache,
	}
}

// RegisterLogin 注册用户
func (b *AuthBiz) RegisterLogin(ctx context.Context, req *auth.RegisterLoginRequest) (*auth.RegisterLoginResponse, error) {
	return &auth.RegisterLoginResponse{}, nil
}

// AuthRefresh 刷新Jwt
func (b *AuthBiz) AuthRefresh(ctx context.Context, req *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	uid, ok := jwtx.FromContextUid(ctx)
	if !ok || uid == 0 {
		return nil, errorx.AuthInvalidError(i18nx.FromContextError(ctx, "AuthInvalid"))
	}
	// 生成JWT
	tokenID := primitive.NewObjectID().Hex()
	token, err := jwtx.GenToken(uid, tokenID)
	if err != nil {
		log.Errorf("GenToken %v", err)
		return nil, errorx.ServiceError(i18nx.FromContextError(ctx, "ServiceError"))
	}
	resp := &auth.RefreshResponse{
		Token: token,
	}
	return resp, nil
}
