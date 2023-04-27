package util

import (
	"github.com/go-kratos/kratos/v2/errors"
	validator "github.com/go-playground/validator/v10"

	"kratos-application/api/gen/api"
)

// ResponseModel 响应请求的公共模型
type ResponseModel struct {
	Code api.ECode `json:"code"`           // 响应码
	Msg  string    `json:"msg"`            // 响应信息
	Data any       `json:"data,omitempty"` // 数据
}

func ResponseError(err error) *ResponseModel {
	if e, ok := err.(*Error); ok {
		return &ResponseModel{
			Code: e.Code,
			Msg:  e.Error(),
		}
	}
	if e, ok := err.(*errors.Error); ok {
		return &ResponseModel{
			Code: api.ECode_EUnknown,
			Msg:  e.Error(),
		}
	}
	if e, ok := err.(validator.ValidationErrors); ok {
		return &ResponseModel{
			Code: api.ECode_EParam,
			Msg:  e.Error(),
		}
	}
	return &ResponseModel{
		Code: api.ECode_EUnknown,
		Msg:  err.Error(),
	}
}
