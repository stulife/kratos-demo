package main

import (
	"flag"
	"fmt"
	nacosconf "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	_ "github.com/go-kratos/kratos/v2/encoding/yaml"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-demo/app/demo/internal/conf"
	"kratos-demo/pkg/nacos"
	conf2 "kratos-demo/pkg/nacos/conf"
	"os"
)

// go build-ldflags "-X main.Version=x.y.z -X main.Name=userService"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server, r registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}
	var bnc conf2.NacosBootstrap
	if err := c.Scan(&bnc); err != nil {
		panic(err)
	}
	nc := config.New(
		config.WithSource(
			nacosconf.NewConfigSource(nacos.NacosConfig(bnc.Nacos), nacosconf.WithGroup(bnc.Nacos.Config.Group), nacosconf.WithDataID(bnc.Nacos.Config.DataId)),
		),
	)

	if err := nc.Load(); err != nil {
		panic(err)
	}
	serviceName, err := nc.Value("service.name").String()
	Version, err = nc.Value("service.version").String()
	id, err = nc.Value("service.id").String()
	Name = serviceName

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)

	var bc conf.Bootstrap
	if err := nc.Scan(&bc); err != nil {
		panic(err)
	}

	if err := nc.Watch("service", func(key string, value config.Value) {
		fmt.Println("config(key=%s) changed: %s\n", key, value.Load())
	}); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, bnc.Nacos, bc.Jwt, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
