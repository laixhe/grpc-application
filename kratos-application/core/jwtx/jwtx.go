package jwtx

import (
	"context"
	"time"

	jwtKratos "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// Config jwt
type Config struct {
	SecretKey  string `mapstructure:"secret_key"`  // jwt secret key
	ExpireTime int    `mapstructure:"expire_time"` // 过期时长(单位秒)
}

// CustomClaims 自定义声明类型
type CustomClaims struct {
	Uid uint32 `json:"uid"`
	jwtv4.RegisteredClaims
}

// GenToken 生成JWT
func GenToken(uid uint32, c *Config) (string, error) {
	claims := CustomClaims{
		Uid: uid,
	}

	nowTime := time.Now()
	// 过期时间
	if c.ExpireTime > 0 {
		claims.ExpiresAt = jwtv4.NewNumericDate(nowTime.Add(time.Duration(c.ExpireTime) * time.Second))
	}
	// 发布时间（创建时间）
	claims.IssuedAt = jwtv4.NewNumericDate(nowTime)
	token := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)
	return token.SignedString([]byte(c.SecretKey))
}

// ParseToken 解析JWT
func ParseToken(tokenString string, c *Config) (*CustomClaims, error) {
	token, err := jwtv4.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtv4.Token) (i interface{}, err error) {
		return []byte(c.SecretKey), nil
	})
	if err != nil {
		ve, ok := err.(*jwtv4.ValidationError)
		if !ok {
			return nil, jwtKratos.ErrTokenInvalid
		}
		if ve.Errors&jwtv4.ValidationErrorMalformed != 0 {
			return nil, jwtKratos.ErrTokenInvalid
		}
		if ve.Errors&(jwtv4.ValidationErrorExpired|jwtv4.ValidationErrorNotValidYet) != 0 {
			return nil, jwtKratos.ErrTokenExpired
		}
		return nil, jwtKratos.ErrTokenParseFail
	}
	if token != nil {
		// 对 token 对象中的 Claim 进行类型断言
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验 token
			return claims, nil
		}
	}
	return nil, jwtKratos.ErrTokenInvalid
}

func FromContextUid(ctx context.Context) (uint32, bool) {
	claims, ok := jwtKratos.FromContext(ctx)
	if ok {
		customClaims, is := claims.(*CustomClaims)
		if is {
			return customClaims.Uid, true
		}
	}
	return 0, false
}
