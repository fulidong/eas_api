package server

import (
	v1 "eas_api/api/eas_api/v1"
	"eas_api/internal/conf"
	"eas_api/internal/middleware"
	"eas_api/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, eas *service.EasService, easSalesPaper *service.EasSalesPaperService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			middleware.RequestIdWithHeader,
			middleware.TryParseHeader(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterEasServiceServer(srv, eas)
	v1.RegisterEasSalesPaperServiceServer(srv, easSalesPaper)
	return srv
}
