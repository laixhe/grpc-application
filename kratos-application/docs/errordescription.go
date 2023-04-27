package docs

import (
	"bytes"
	"os"
)

// ErrorDescription api文档错误码
func ErrorDescription() string {
	data, err := os.ReadFile("./api/error.proto")
	if err != nil {
		return ""
	}
	splits := bytes.Split(data, []byte("enum"))
	if len(splits) <= 1 {
		return ""
	}
	data = bytes.Join(splits[1:], []byte{})
	return "enum" + string(data)
}
