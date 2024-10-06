package requestx

import (
	"context"

	"kratos-application/api/gen/enum/eapp"
)

// 头部 IP

type IPContextKey struct{}

const IPHeaderKey = "ip"

// 头部 平台

type PlatformContextKey struct{}

const PlatformHeaderKey = "platform"

func GetPlatform(ctx context.Context) string {
	platform, ok := ctx.Value(PlatformContextKey{}).(string)
	if ok {
		value := eapp.Platform_value[platform]
		if value != 0 {
			return platform
		}
	}
	return eapp.Platform_unknown.String()
}
