package server

import (
	"net/http"
	"strings"

	kratosHttp "github.com/go-kratos/kratos/v2/transport/http"
)

// ResponseEncoder http 正确-返回统一结构
func ResponseEncoder() kratosHttp.EncodeResponseFunc {
	return func(w http.ResponseWriter, r *http.Request, v interface{}) error {
		if v == nil {
			return nil
		}
		if rd, ok := v.(kratosHttp.Redirector); ok {
			url, code := rd.Redirect()
			http.Redirect(w, r, url, code)
			return nil
		}
		codec, _ := kratosHttp.CodecForRequest(r, "Accept")

		// 统一结构
		data := []byte{}

		w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
		_, err := w.Write(data)
		if err != nil {
			return err
		}
		return nil
	}
}

// ErrorEncoder http 错误-返回统一结构
func ErrorEncoder() kratosHttp.EncodeErrorFunc {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		codec, _ := kratosHttp.CodecForRequest(r, "Accept")

		// 统一结构
		body := []byte{}

		w.Header().Set("Content-Type", strings.Join([]string{"application", codec.Name()}, "/"))
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(body)
	}
}
