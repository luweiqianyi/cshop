---
sidebar: auto
prev: false
next: /cshop/01-auth-rpc.html
---
# 项目基本环境搭建
## 项目初始化
```
go mod init cshop
```
## go-zero
`go-zero`微服务框架获取。
```
go get -u github.com/zeromicro/go-zero@latest
```
## goctl脚手架
获取goctl脚手架
```
go get -u github.com/zeromicro/go-zero/tools/goctl@latest
```
## grpc
获取`grpc`相关的组件: `protoc`,`protoc-gen-go`,`protoc-gen-go-grpc`
```
goctl env check --install --verbose --force
```
成功输出如下：
```
[goctl-env]: preparing to check env

[goctl-env]: looking up "protoc"
[goctl-env]: "protoc" is installed

[goctl-env]: looking up "protoc-gen-go"
[goctl-env]: "protoc-gen-go" is installed

[goctl-env]: looking up "protoc-gen-go-grpc"
[goctl-env]: "protoc-gen-go-grpc" is installed

[goctl-env]: congratulations! your goctl environment is ready!
```

