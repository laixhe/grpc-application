package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-playground/validator/v10"
	"kratos-application/util"
)

// 参数验证
var validate = validator.New()

// ValidateMiddleware 参数验证
func ValidateMiddleware() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if err := validate.Struct(req); err != nil {
				return nil, util.ParamError(err)
			}
			return handler(ctx, req)
		}
	}
}
