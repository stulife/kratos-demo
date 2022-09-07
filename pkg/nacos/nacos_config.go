package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"kratos-demo/pkg/nacos/conf"
)

func NacosConfig(conf *conf.Nacos) config_client.IConfigClient {

	sc := []constant.ServerConfig{
		*constant.NewServerConfig(conf.Config.Ip, conf.Config.Port),
	}
	cc := &constant.ClientConfig{
		NamespaceId:         conf.Config.NamespaceId,
		TimeoutMs:           conf.Config.TimeoutMs,
		NotLoadCacheAtStart: conf.Config.NotLoadCacheAtStart,
		LogDir:              conf.Config.LogDir,
		CacheDir:            conf.Config.CacheDir,
		LogLevel:            conf.Config.LogLevel,
	}

	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	return client
}
