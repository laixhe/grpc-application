package gormx

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/logger"
)

type Writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *Writer {
	return &Writer{Writer: w}
}

// Printf 格式化打印日志
func (w *Writer) Printf(message string, data ...interface{}) {
	log.Info(fmt.Sprintf(message, data...))
}
