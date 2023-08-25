---
sidebar: auto
prev: /cshop/02-auth-api.html
next: /cshop/03-account-api-02.html
---
# account-api(一)
本教程演示从零到一演示如何创建一个`api`服务。本过程以用户账号的注册、注销、登录、退出这四个过程来演示。

## 首先确实前端请求的格式
### 注册
* 请求方法：`post`
* 请求路径：`/register`
* 请求参数
    * `accountName`: 账号名
    * `password`: 账号密码

### 注销
* 请求方法：`get`
* 请求路径：`/unregister`
* 请求参数
    * `accountName`: 账号名

### 登录
* 请求方法：`post`
* 请求路径：`/login`
* 请求参数
    * `accountName`: 账号名
    * `password`: 账号密码

### 退出登录
* 请求方法：`get`
* 请求路径：`/logout`
* 请求参数
    * `accountName`: 账号名

## 根据以上请求格式，定义`account.api`文件
`account.api`文件保存路径：`cmd/account/api/`
文件内容如下:
```
syntax = "v1"

info(
    title: "account-api"
    desc: "account-api"
    author: "luweiqianyi"
    email: "runningriven@gmail.com"
    version: "1.0.0"
)

type CommonResp {
    Success bool   `json:"success"`
    Detail  string `json:"detail,omitempty"`
}

type (
    RegisterReq {
        AccountName string `json:"accountName"`
        Password    string `json:"password"`
    }
    RegisterResp {
        CommonResp
    }

    UnRegisterReq {
        AccountName string `json:"accountName"`
    }
    UnRegisterResp {
        CommonResp
    }
)


type (
    LoginReq {
        AccountName string `json:"accountName"`
        Password    string `json:"password"`
    }

    LoginResp {
        CommonResp
        Token string `json:"token"`
    }

    LogoutReq {
        AccountName string `json:"accountName"`
    }

    LogoutResp {
        CommonResp
    }
)


service account-api {
    @handler RegisterHandler
    post /register (RegisterReq) returns (RegisterResp)

    @handler UnRegisterHandler
    get /unregister (UnRegisterReq) returns (UnRegisterResp)

    @handler LoginHandler
    post /login(LoginReq) returns (LoginResp)

    @handler LogoutHandler
    get /logout(LogoutReq) returns (LogoutResp)
}
```
## 新建脚本文件，用于自动生成api代码
在目录`cmd/account/api/`下新建脚本文件`auto-gen-api-code.bat`。脚本内增加：
```shell
goctl api go -api account.api -dir . -style gozero
```
## 执行上述脚本文件
## 修改代码，使其能够运行
1. 修改自动生成的`account-api.yaml`文件，将`Port`改成`9003`
2. 修改`account.go`文件，修改一下配置文件的读取路径，修改如下：
```go
    path, _ := os.Executable()
	dir := filepath.Dir(path)
	fullPath := filepath.Join(dir, "./etc/auth-api.yaml")
	flag.Set("f", fullPath)
```
3. 修改`cmd/account/api/internal/logic`目录下的四个`*logic.go`文件，这四个文件对应`account.api`文件中定义的四个接口，这里简单一点，对于四个接口都直接返回正常响应即可。下面以`/login`接口为例，给出代码示例：
```go
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
    // 这里不验证，直接给个成功返回即可
	resp = new(types.LoginResp)
	resp.CommonResp.Success = true
	return
}
```
4. 运行起来

## 编写测试程序
`cmd/account/api/test`目录下新建`api_test.go`,测试代码如下(先展示一个`/login`接口)：
```go
package test

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"testing"
)

const (
	auth_api_host = "127.0.0.1"
	port          = "9003"
	register      = "/register"
	unregister    = "/unregister"
	login         = "/login"
	logout        = "/logout"

	accountNameKey = "accountName"
	passwordKey    = "password"
)

func TestLogin(t *testing.T) {
	url := fmt.Sprintf("http://%v:%v%v", auth_api_host, port, login)

	resp, err := resty.New().R().
		SetHeader("Access-Token", token).
		SetBody(map[string]interface{}{accountNameKey: "leebai", passwordKey: "123456"}).
		Post(url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("resp: %v\n", resp)
}
```
## 运行，模拟一次登录请求发送，查看请求结果
最终结果如下所示：
```log
=== RUN   TestLogin
resp: {"success":true,"token":""}
--- PASS: TestLogin (0.00s)
PASS

Process finished with the exit code 0
```

以上响应返回了`"success"`和`"token"`，这和我们上面定义的`account.api`中对于`LoginResp`的定义相符合。其他的三个接口，类比可实现。

至此，就说明初步完成了`accout-api`服务的功能。