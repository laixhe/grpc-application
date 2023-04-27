package util

import "kratos-application/api/gen/api"

type Error struct {
	Code api.ECode
	Err  error
}

func (e *Error) Error() string {
	if e.Err == nil {
		return e.Code.String()
	}
	return e.Err.Error()
}

func NewError(code api.ECode, err error) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

func UnknownError(err error) *Error {
	return &Error{
		Code: api.ECode_EUnknown,
		Err:  err,
	}
}

func ServiceError(err error) *Error {
	return &Error{
		Code: api.ECode_EService,
		Err:  err,
	}
}

func ParamError(err error) *Error {
	return &Error{
		Code: api.ECode_EParam,
		Err:  err,
	}
}
