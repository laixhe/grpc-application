package i18nx

import (
	"context"
	"errors"

	"github.com/gogf/gf/v2/i18n/gi18n"
)

// 头部

type LanguageContextKey struct{}

const LanguageHeaderKey = "accept-language"

var (
	i18n = gi18n.New()

	CtxZhCn = context.Background()
	CtxEn   = context.Background()
)

func Init() {
	i18n.SetLanguage(LanguageZhCn)
	CtxZhCn = gi18n.WithLanguage(context.Background(), LanguageZhCn)
	CtxEn = gi18n.WithLanguage(context.Background(), LanguageEn)
}

func FromContext(ctx context.Context, str string) string {
	language, ok := ctx.Value(LanguageContextKey{}).(string)
	if !ok {
		return i18n.Translate(CtxZhCn, str)
	}
	if language == LanguageEn {
		return i18n.Translate(CtxEn, str)
	}
	return i18n.Translate(CtxZhCn, str)
}

func FromContextError(ctx context.Context, str string) error {
	errStr := FromContext(ctx, str)
	return errors.New(errStr)
}
