package errorx

import (
	"kratos-application/api/gen/ecode"
)

type Error struct {
	Code ecode.ECode
	Err  error
}

func (e *Error) Error() string {
	if e.Err == nil {
		return e.Code.String()
	}
	return e.Err.Error()
}

func NewError(code ecode.ECode, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

func ServiceError(err error) *Error {
	return &Error{
		Code: ecode.ECode_Service,
		Err:  err,
	}
}

func ParamError(err error) *Error {
	return &Error{
		Code: ecode.ECode_Param,
		Err:  err,
	}
}

func TipMessageError(err error) *Error {
	return &Error{
		Code: ecode.ECode_TipMessage,
		Err:  err,
	}
}

func AuthInvalidError(err error) *Error {
	return &Error{
		Code: ecode.ECode_AuthInvalid,
		Err:  err,
	}
}

func UserNotExistError(err error) *Error {
	return &Error{
		Code: ecode.ECode_UserNotExist,
		Err:  err,
	}
}
