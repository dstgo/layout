package app

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	example "github.com/dstgo/layout/api/gen/example/v1"
	"github.com/dstgo/layout/internal/conf"
)

type Example struct {
	example.UnimplementedExampleServer
}

func NewExample(c *conf.Config) *Example {
	return &Example{}
}

func init() {
	register(func(c *conf.Config, hs *http.Server, gs *grpc.Server) {
		srv := NewExample(c)
		example.RegisterExampleServer(gs, srv)
		example.RegisterExampleHTTPServer(hs, srv)
	})
}
