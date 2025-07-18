package iheader

import (
	"strings"

	"github.com/go-kratos/kratos/v2/transport"
)

const (
	ResponseContentJsonType = "application/json" // json 数据
	ResponseContentTextType = "text/plain"       // 文本数据
	TokenHeaderKey          = "X-Token"          // 用户token
	ForwardForHeaderKey     = "X-Forwarded-For"  // 客户端ip
	RequestIdKey            = "X-Request-Id"     // request_id
)

func GetToken(h transport.Header) string {
	return h.Get(TokenHeaderKey)
}

func GetClientIp(h transport.Header) string {
	value := h.Get(ForwardForHeaderKey)
	splits := strings.Split(value, ",")
	if len(splits) > 0 {
		return splits[0]
	}
	return ""
}

func GetRequestId(h transport.Header) string {
	return h.Get(RequestIdKey)
}
