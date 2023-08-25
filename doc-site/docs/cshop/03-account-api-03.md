---
sidebar: auto
prev: /cshop/03-account-api-02.html
next: false
---
# account-api服务集成其他rpc服务
## 前景
在`account-api`服务的登录接口中有一段逻辑，当用户请求中的账号和密码经过服务器的验证通过之后，服务器需要向用户返回一个`token`,作为用户后续访问其他服务的身份认证。

在这里，这个`token`的生成管理不应该由我们的`account-api`服务来管理，在我们之前的文章中，将这部分`token`生成的逻辑以及`token`验证的都放到`auth-rpc`服务里面去了。接下来就主要介绍如何在我们这个`account-api`服务中集成对`auth-rpc`服务的访问即可。

## 实现过程
### 确定远程`auth-rpc`服务的访问地址
这里是`127.0.0.1:9000`,所以在配置文件`account-api.yaml`中配置好远程`auth-rpc`服务的访问地址，如下所示:
```yaml
AuthRpcClientConf:
  Endpoints:
    - 127.0.0.1:9000
```
### 修改配置代码
修改`cmd/account/api/internal/config/config.go`,如下所示：
```go
type Config struct {
	rest.RestConf

	MySQL struct {
		DataSource string
	}

	Salt string

	AuthRpcClientConf zrpc.RpcClientConf
}
```
### 创建rpc客户端连接对象
```go
type ServiceContext struct {
	Config             config.Config
	TbUserAccountModel model.TbUserAccountModel
	AuthRpcClient      auth.Auth // 连接远程rpc的客户端对象
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	return &ServiceContext{
		Config:             c,
		TbUserAccountModel: model.NewTbUserAccountModel(conn),
		AuthRpcClient:      auth.NewAuth(zrpc.MustNewClient(c.AuthRpcClientConf)),
	}
}
```

### 修改登录逻辑，增加token相关处理
具体代码详见`cmd/account/api/internal/logic/loginlogic.go`。

### 测试过程
1. 先启动`auth-rpc`服务
2. 再启动`account-api`服务
3. 测试程序测试, 测试程序位于`cmd/accout/api/test/api_test.go`
4. 输出结果如下
```log
=== RUN   TestLogin
resp: {"success":true,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiYWNjb3VudE5hbWVcIjpcImxlZWJhaVwifSIsImV4cCI6MTY5NDIzMDkyNH0.yWpSxY5yfzMOQunlP7uQDrhbZh8mz5A8pDOATDcpjEI"}
--- PASS: TestLogin (0.23s)
PASS
```
> 如上所示：token成功生成并返回了。目前存在的不足之处时，`account-api`服务的启动依赖于`auth-rpc`服务的启动，后者不启动，前者启动不了，后期再考虑优化。