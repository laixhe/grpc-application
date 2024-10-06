package responsex

import (
	"errors"
	"strings"

	kratosErrors "github.com/go-kratos/kratos/v2/errors"

	"kratos-application/api/gen/ecode"
	"kratos-application/core/errorx"
	"kratos-application/core/jwtx"
)

// ResponseModel 响应请求的公共模型
type ResponseModel struct {
	Code ecode.ECode `json:"code"`           // 响应码
	Msg  string      `json:"msg"`            // 响应信息
	Data any         `json:"data,omitempty"` // 数据
}

func ResponseError(err error) *ResponseModel {
	var ex *errorx.Error
	if errors.As(err, &ex) {
		return &ResponseModel{
			Code: ex.Code,
			Msg:  ex.Error(),
		}
	}
	var e *kratosErrors.Error
	if errors.As(err, &e) {
		eCode := ecode.ECode_Service
		eMsg := e.GetMessage()
		if jwtCode, jwtIs := jwtx.ErrorToCode(err); jwtIs {
			eCode = jwtCode
		}
		if eCode == ecode.ECode_Service {
			if strings.HasPrefix(e.GetMessage(), "body unmarshal") ||
				strings.HasPrefix(e.GetMessage(), "unregister Content-Type") {
				eCode = ecode.ECode_Param
				eMsg = "param error: " + eMsg
			}
		}
		return &ResponseModel{
			Code: eCode,
			Msg:  eMsg,
		}
	}
	return &ResponseModel{
		Code: ecode.ECode_Service,
		Msg:  err.Error(),
	}
}
