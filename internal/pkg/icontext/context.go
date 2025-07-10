package icontext

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"os"
)

const (
	clientIP     = "X-Forwarded-For" // 客户端IP
	requestIdKey = "RequestId"       // req id
	userIdKey    = "userIdKey"       // 用户ID
	userNameKey  = "userNameKey"     // 用户名称
	userRuleKey  = "userRuleKey"     // 用户角色
)

func withValue(ctx context.Context, key, value string) context.Context {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		md = metadata.Metadata{}
	}
	md.Set(key, value)
	return metadata.NewServerContext(ctx, md)
	//return metadata.AppendToClientContext(ctx, key, value)
}

func fromValue(ctx context.Context, key string) (string, bool) {
	md, ok := metadata.FromServerContext(ctx)
	if !ok {
		return "", false
	}

	out := md.Get(key)
	return out, out != ""
}

// 客户端ip
func WithClientIP(ctx context.Context, in string) context.Context {
	return withValue(ctx, clientIP, in)
}

func ClientIPFrom(ctx context.Context) (string, bool) {
	return fromValue(ctx, clientIP)
}

func WithUserIdKey(ctx context.Context, in string) context.Context {
	return withValue(ctx, userIdKey, in)
}

func UserIdFrom(ctx context.Context) (string, bool) {
	return fromValue(ctx, userIdKey)
}

func WithUserNameKey(ctx context.Context, in string) context.Context {
	return withValue(ctx, userNameKey, in)
}

func UserNameFrom(ctx context.Context) (string, bool) {
	return fromValue(ctx, userNameKey)
}

func WithUserRuleKey(ctx context.Context, in string) context.Context {
	return withValue(ctx, userRuleKey, in)
}

func UserRuleFrom(ctx context.Context) (string, bool) {
	return fromValue(ctx, userRuleKey)
}

// request id

func WithRequestId(ctx context.Context, in string) context.Context {
	return withValue(ctx, requestIdKey, in)
}

func RequestIdFrom(ctx context.Context) (string, bool) {
	return fromValue(ctx, requestIdKey)
}

// context

func LoggerValues() []interface{} {
	return []interface{}{

		"client_ip", log.Valuer(func(ctx context.Context) interface{} {
			clientIp, _ := ClientIPFrom(ctx)
			return clientIp
		}),
		"user_id", log.Valuer(func(ctx context.Context) interface{} {
			userId, _ := UserIdFrom(ctx)
			return userId
		}),
		"request_id", log.Valuer(func(ctx context.Context) interface{} {
			reqId, _ := RequestIdFrom(ctx)
			return reqId
		}),
		"namespace", log.Valuer(func(ctx context.Context) interface{} {
			return os.Getenv("NAMESPACE")
		}),
	}
}
