package server

import (
	"encoding/json"
	stdHttp "net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/transport/http"

	"kratos-application/util"
)

// ResponseEncoder http 正确-返回统一结构
func ResponseEncoder() http.EncodeResponseFunc {
	return func(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
		if v == nil {
			return nil
		}
		if rd, ok := v.(http.Redirector); ok {
			url, code := rd.Redirect()
			stdHttp.Redirect(w, r, url, code)
			return nil
		}
		codec, _ := http.CodecForRequest(r, "Accept")

		// 统一结构
		data, err := json.Marshal(util.ResponseModel{Data: v})
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
		_, err = w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}

// ErrorEncoder http 错误-返回统一结构
func ErrorEncoder() http.EncodeErrorFunc {
	return func(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
		codec, _ := http.CodecForRequest(r, "Accept")

		// 统一结构
		body, err := json.Marshal(util.ResponseError(err))
		if err != nil {
			w.WriteHeader(stdHttp.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
		w.WriteHeader(stdHttp.StatusOK)
		_, _ = w.Write(body)
	}
}
