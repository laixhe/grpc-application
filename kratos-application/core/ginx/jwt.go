package jwtx

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"kratos-application/core/jwtx"
	"kratos-application/core/responsex"
)

// JwtAuth 鉴权
func JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var parseTokenErr error
		token := ctx.Request.Header.Get(jwtx.Authorization)
		if len(token) > 0 {
			if strings.HasPrefix(token, jwtx.Bearer) {
				claims, err := jwtx.ParseToken(token[jwtx.BearerLen:])
				if err == nil {
					ctx.Set(jwtx.AuthorizationClaimsHeaderKey, claims)
					ctx.Next()
					return
				}
				parseTokenErr = err
			}
		}
		ctx.JSON(http.StatusOK, responsex.ResponseError(parseTokenErr))
		// 返回错误
		ctx.Abort()
	}
}

func GinContextUid(c *gin.Context) (uint64, bool) {
	value, exists := c.Get(jwtx.AuthorizationClaimsHeaderKey)
	if exists {
		customClaims, is := value.(*jwtx.CustomClaims)
		if is {
			return customClaims.Uid, true
		}
	}
	return 0, false
}
