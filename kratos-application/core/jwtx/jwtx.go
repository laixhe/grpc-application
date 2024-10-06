package jwtx

import (
	"context"
	"errors"
	"time"

	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv5 "github.com/golang-jwt/jwt/v5"

	"kratos-application/api/gen/config/cauth"
	"kratos-application/api/gen/ecode"
	"kratos-application/core/conf"
)

// Authorization
const (
	Authorization = "Authorization"
	Bearer        = "Bearer "
	BearerLen     = 7
)

type AuthorizationContextKey struct{}

type IsAuthorizationContextKey struct{}

const AuthorizationClaimsHeaderKey = "AuthorizationClaims"

// CustomClaims 自定义声明类型 并内嵌 jwt.RegisteredClaims
// jwt 包自带的 jwt.RegisteredClaims 只包含了官方字段
type CustomClaims struct {
	Uid uint64 `json:"uid"`
	jwtv5.RegisteredClaims
}

// Checking 检查
func Checking(c *cauth.Jwt) {
	if c == nil {
		panic(errors.New("config jwt is nil"))
	}
	if c.SecretKey == "" {
		panic(errors.New("config jwt secret_key is empty"))
	}
	if c.GetExpireTime() == nil {
		panic(errors.New("config jwt expire_time is empty"))
	}
}

// GenToken 生成JWT
func GenToken(uid uint64, id string) (string, error) {
	claims := CustomClaims{
		Uid: uid,
		RegisteredClaims: jwtv5.RegisteredClaims{
			ID: id,
		},
	}
	nowTime := time.Now()
	// 过期时间
	if conf.Get().GetJwt().GetExpireTime().IsValid() {
		if conf.Get().GetJwt().GetExpireTime().Seconds > 0 || conf.Get().GetJwt().GetExpireTime().Nanos > 0 {
			claims.ExpiresAt = jwtv5.NewNumericDate(nowTime.Add(conf.Get().GetJwt().GetExpireTime().AsDuration()))
		}
	}
	// 发布时间（创建时间）
	claims.IssuedAt = jwtv5.NewNumericDate(nowTime)
	// 生效时间
	claims.NotBefore = jwtv5.NewNumericDate(nowTime)
	// 使用指定的签名方法创建签名对象（使用 HS256 签名算法）
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	// 使用指定的 secret 签名并获得完整的编码后的字符串 token
	return token.SignedString([]byte(conf.Get().GetJwt().GetSecretKey()))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwtv5.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtv5.Token) (i interface{}, err error) {
		return []byte(conf.Get().GetJwt().GetSecretKey()), nil
	})
	if err != nil {
		if errors.Is(err, jwtv5.ErrTokenExpired) {
			return nil, jwt.ErrTokenExpired
		}
		return nil, jwt.ErrTokenInvalid
	}
	if token != nil {
		// 对 token 对象中的 Claim 进行类型断言
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验 token
			return claims, nil
		}
	}
	return nil, jwt.ErrTokenInvalid
}

func ErrorToCode(err error) (ecode.ECode, bool) {
	if errors.Is(err, jwt.ErrMissingJwtToken) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrMissingKeyFunc) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrTokenInvalid) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		return ecode.ECode_AuthExpire, true
	}
	if errors.Is(err, jwt.ErrTokenParseFail) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrUnSupportSigningMethod) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrWrongContext) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrNeedTokenProvider) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrSignToken) {
		return ecode.ECode_AuthInvalid, true
	}
	if errors.Is(err, jwt.ErrGetKey) {
		return ecode.ECode_AuthInvalid, true
	}
	return ecode.ECode_AuthInvalid, false
}

func FromContextUid(ctx context.Context) (uint64, bool) {
	claims, ok := jwt.FromContext(ctx)
	if ok {
		customClaims, is := claims.(*CustomClaims)
		if is {
			return customClaims.Uid, true
		}
	}
	return 0, false
}

func FromContextID(ctx context.Context) (string, bool) {
	claims, ok := jwt.FromContext(ctx)
	if ok {
		customClaims, is := claims.(*CustomClaims)
		if is {
			return customClaims.ID, true
		}
	}
	return "", false
}
