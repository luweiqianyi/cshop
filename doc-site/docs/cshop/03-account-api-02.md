---
sidebar: auto
prev: /cshop/03-account-api-01.html
next: false
---
# account-api集成mysql
本文档主要介绍如何在`account-api`服务中引入`mysql`。

## 详细过程
### 确定数据来源
`127.0.0.1:3306/go_zero_demo`的数据库中之前已经存储了相应的数据。所以后续我们用`goctl model`来自动生成相关代码来访问该数据库。

### 自动生成数据库访问代码
在目录`/cmd/account/api`目录下面执行`goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/go_zero_demo" -table="*"  -dir="./model"`

控制台(Terminal)输出`Done`说明自动生成数据库访问代码成功。同时，目录`/cmd/account/api`下面多了一个`model`的目录，里面存放就是该命令生成的代码。

### 使用生成的代码访问数据库(CRUD)
1. 修改`account-api.yaml`，增加数据库配置
    ```yaml
    MySQL:
        DataSource: root:123456@tcp(127.0.0.1:3306)/go_zero_demo?charset=utf8mb4&parseTime=true
    ```
2. 修改`cmd/account/api/internal/config/config.go`，增加配置
    ```go
    type Config struct {
        rest.RestConf

        MySQL struct {
            DataSource string
        }
    }
    ```
3. 修改`cmd/account/api/internal/svc/servicecontext.go`,添加对某个表访问的映射对象
    ```go
    type ServiceContext struct {
        Config             config.Config
        TbUserAccountModel model.TbUserAccountModel
    }

    func NewServiceContext(c config.Config) *ServiceContext {
        conn := sqlx.NewMysql(c.MySQL.DataSource)
        return &ServiceContext{
            Config:             c,
            TbUserAccountModel: model.NewTbUserAccountModel(conn),
        }
    }
    ```
4. 以注册接口为例，给出注册过程的实现逻辑
    ```go
    func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
        record := &model.TbUserAccount{
            AccountName: sql.NullString{
                String: req.AccountName,
                Valid:  true,
            },
            Password: sql.NullString{
                String: req.Password,
                Valid:  true,
            },
        }

        _, err = l.svcCtx.TbUserAccountModel.Insert(l.ctx, record)
        if err != nil {
            return &types.RegisterResp{
                CommonResp: types.CommonResp{
                    Success: false,
                },
            }, fmt.Errorf("register failed, err: %v", err)
        }

        return &types.RegisterResp{
            CommonResp: types.CommonResp{
                Success: true,
            },
        }, nil
    }
    ```
### 编译，生成并运行`account-api`服务。
### 测试。
之前在我们的`cmd/account/api/test/api_test.go`中已经写好了的客户端代码。测试即可。
运行结果
* 成功：
    ```log
    === RUN   TestRegister
    resp: {"success":true}
    --- PASS: TestRegister (0.13s)
    PASS
    ```

* 失败：
    ```log
    === RUN   TestRegister
    resp: register failed, err: Error 1062 (23000): Duplicate entry 'leebai' for key 'TbUserAccount.accountName'
    --- PASS: TestRegister (0.18s)
    PASS
    ```
### 至此，便完成了从前端请求到后端将数据存储到mysql的整个过程。