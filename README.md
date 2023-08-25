# cshop
本项目是一个基于[go-zero](https://go-zero.dev/)开发的一个项目。

## 快速开始
1. 克隆项目
    ```git
    git clone https://github.com/luweiqianyi/cshop.git
    ```
2. 进入`cmd`的各服务模块，启动
3. 进入各个服务模块的`test`目录，运行对应的测试脚本函数即可

## 文档
克隆项目后,进入`doc-site`目录，控制台运行命令`npm run docs:dev`即可在本地建立了本地可访问的站点，通过`localhost:8080`即可访问。该站点主要是对本项目的一些实现过程简要介绍。

## 目录介绍
* `cmd`: 各个微服务模块的源码目录。
* `docs`: 本项目的文档目录，也是部署文档站点的根目录。
* `docker-env`: 各个服务的容器环境的挂载目录，以及项目部署文件存放的目录。
* `pkg`: 各个微服务模块引用的公共代码的目录。

## 功能模块
### auth
鉴权服务。为项目的其他服务提供一个统一的鉴权功能。
#### auth-api
鉴权服务api接口。
#### auth-rpc
鉴权服务rpc接口。

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
