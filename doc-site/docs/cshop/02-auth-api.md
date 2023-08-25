---
sidebar: auto
prev: /cshop/01-auth-rpc.html
next: /cshop/03-account-api-01.html
---
# auth-api
`auth-api`服务，为其他服务提供基于`http`的鉴权服务。

以下介绍如何定义并启动一个`auth-api`服务。

## auth.api接口定义
```api
syntax = "v1"

info(
    title: "auth-api"
    desc: "auth-api"
    author: "luweiqianyi"
    email: "runningriven@gmail.com"
    version: "1.0.0"
)

type CommonResp {
    Success  bool `json:"success"`
    Detail string `json:"detail,omitempty"`
}

type (
    AuthReq {
        AccessToken string `header:Access-Token`
    }

    AuthResp {
        CommonResp
    }
)

service auth-api {
    @handler AuthHandler
    get /auth(AuthReq) returns(AuthResp)
}
```
* 上面定义了一个`get`方式的`/auth`请求。用于用户身份的鉴权。
## 自动生成api相关代码
1. 进入`cmd/auth/api`目录，新建`auto-gen-api-code.bat`脚本文件
2. 在上述的脚本文件中添加: `goctl api go -api auth.api -dir . -style gozero`
3. 执行上述脚本文件
4. 执行成功会自动生成`api`相关代码

## 添加自定义验证逻辑
修改`cmd/auth/api/internal/logic`目录下的`authlogic.go`文件，添加验证逻辑。
```go
func (l *AuthLogic) Auth(req *types.AuthReq) (resp *types.AuthResp, err error) {
	resp = new(types.AuthResp)

	_, err = l.svcCtx.AuthRpcCli.ValidateToken(context.Background(), &pb.TokenValidateReq{
		Token: req.AccessToken,
	})
	if err != nil {
		resp.CommonResp.Success = false
		resp.CommonResp.Detail = fmt.Sprintf("auth failed: %v", err)
		return
	}

	resp.CommonResp.Success = true
	return
}
```

其他配置相关代码想看与本文件相同时间提交的相关代码。
## 编写测试程序
`cmd/auth/api/test`目录下新建`api_test.go`文件，测试代码如下所示。
```go
package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
)

const (
	auth_api_host = "127.0.0.1"
	port          = "9001"

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoie1wiYWNjb3VudE5hbWVcIjpcImxlZWJhaVwifSIsImV4cCI6MTY5NDA4MDA0Nn0.uzPqBT-de3nb1OhKDWqO5XUzGYdvwgKl4qCj4SCLKJQ"
)

func TestSendAuthRequest(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v/auth", auth_api_host, port)

	resp, err := resty.New().R().
		SetHeader("Access-Token", token).
		Get(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}
```
## 验证
按照顺序执行以下过程。
* 启动`auth-rpc`服务
* 启动`auth-api`服务
* 启动上面的测试程序，即可看到运行结果。比如如下所示：
```log
=== RUN   TestSendAuthRequest
resp: {"success":true}
--- PASS: TestSendAuthRequest (0.01s)
PASS

Process finished with the exit code 0
```
