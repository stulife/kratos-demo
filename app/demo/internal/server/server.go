package server

import (
	"github.com/google/wire"
	"kratos-demo/pkg/nacos"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, nacos.NacosConfig, nacos.NewRegistrar)
