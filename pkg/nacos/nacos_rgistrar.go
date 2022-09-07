package nacos

import (
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"kratos-demo/pkg/nacos/conf"
)

// NewRegistrar 服务注册业务注入
func NewRegistrar(conf *conf.Nacos) registry.Registrar {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Discovery.Ip, conf.Discovery.Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         conf.Discovery.NamespaceId,
		TimeoutMs:           conf.Discovery.TimeoutMs,
		NotLoadCacheAtStart: conf.Discovery.NotLoadCacheAtStart,
		LogDir:              conf.Discovery.LogDir,
		CacheDir:            conf.Discovery.CacheDir,
		LogLevel:            conf.Discovery.LogLevel,
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	return nacos.New(client, nacos.WithGroup(conf.Discovery.Group))
}
