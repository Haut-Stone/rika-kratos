package server

import (
	"github.com/felixge/fgprof"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/pprof"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, video *service.VideoService, demo *service.DemoService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/debug/pprof/", pprof.NewHandler())
	srv.Handle("/debug/fgprof", fgprof.Handler())
	v1.RegisterGreeterHTTPServer(srv, greeter)
	v1.RegisterVideoHTTPServer(srv, video) // ! 这里要手动注册,只前的东西才会加载
	v1.RegisterDemoHTTPServer(srv, demo)
	return srv
}
