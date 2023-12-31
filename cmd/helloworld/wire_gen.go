// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"helloworld/internal/conf"
	"helloworld/internal/data"
	"helloworld/internal/server"
	"helloworld/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	videoRepo := data.NewVideoRepo(dataData, logger)
	videoUsecase := biz.NewVideoUsecase(videoRepo, logger)
	videoService := service.NewVideoService(videoUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, videoService, logger)
	demoRepo := data.NewDemoRepo(dataData, logger)
	demoUsecase := biz.NewDemoUsecase(demoRepo, logger)
	demoService := service.NewDemoService(demoUsecase)
	httpServer := server.NewHTTPServer(confServer, greeterService, videoService, demoService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
