# cshop
本项目是一个基于[go-zero](https://go-zero.dev/)开发的一个项目。

## 目录介绍
* `cmd`: 各个微服务模块的源码目录。
* `docs`: 本项目的文档目录，也是部署文档站点的根目录。
* `docker-env`: 各个服务的容器环境的挂载目录，以及项目部署文件存放的目录。
* `pkg`: 各个微服务模块引用的公共代码的目录。

## 功能模块
### gateway
网关服务。为项目的其他服务提供一个统一的鉴权功能。
#### gateway-api
网关服务api接口。
#### gateway-rpc
网关服务rpc接口。

### account
账号服务。平台用户账号管理。
#### account-api
账号服务api接口。
#### account-rpc
账号服务rpc接口。

### user
用户信息服务。平台用户私人信息管理。
#### user-api
用户信息服务api接口。
#### user-rpc
用户信息服务rpc接口。