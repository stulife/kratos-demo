# Kratos Project Template

## Kratos 脚手架
```
集成 Swagger
nacos 配制中心
nacos 服务注册
数据库 基于GORM
auth jwt认证 selector实现白名单
token 基于redis 实现过期、自动续期
CORS 跨域
参数校验 Validate 
上下文对像
自定义http 输出结构
自定义 业务码处理
提供 验证码接口、登录接口,创建用户、获取登录用户信息

```
##  持续更新......
```
集成gateway 
路由与负载均衡
熔断
Metrics 
限流
Tracing 
casbin 
```
##  nacos 配制
```	
kratos_use_config.yaml

service:
  name: kratos_use
  version: 0.0.1
  id: 101
server:
  http:
    addr: 0.0.0.0:8000
    timeout: 1s
  grpc:
    addr: 0.0.0.0:9000
    timeout: 1s
data:
  database:
    driver: mysql
    source: root:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local
  redis:
    addr: 127.0.0.1:6379
    password: a123456
    db: 0
    dial_timeout: 5s
    read_timeout: 5s
    write_timeout: 5s
jwt:
  header: Authorization
  secret: demo
  expire_time: 3600s
```
## sql
```
CREATE TABLE `tb_user` (
`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
`user_name` varchar(50) DEFAULT NULL,
`mobile` varchar(20) DEFAULT NULL,
`nickname` varchar(40) DEFAULT NULL,
`password` varchar(80) DEFAULT NULL,
`status` tinyint(4) DEFAULT NULL,
`email` varchar(40) DEFAULT NULL,
`sex` int(11) DEFAULT NULL,
`avatar` varchar(120) DEFAULT NULL,
`remark` varchar(255) DEFAULT NULL,
`is_admin` tinyint(4) DEFAULT NULL,
`address` varchar(120) DEFAULT NULL,
`created_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```
## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

