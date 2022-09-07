//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"kratos-demo/app/demo/internal/biz"
	"kratos-demo/app/demo/internal/conf"
	"kratos-demo/app/demo/internal/data"
	"kratos-demo/app/demo/internal/server"
	"kratos-demo/app/demo/internal/service"
	nconf "kratos-demo/pkg/nacos/conf"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *nconf.Nacos, *conf.Jwt, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
