---
sidebar: auto
---
# auth-rpc
1. `auth.proto`定义`rpc`接口
    ```text
    syntax = "proto3";

    option go_package = "./pb";

    package pb;

    message TokenValidateReq{
        string token=1;
    }

    message TokenValidateResp{
        bool ok=1;
    }

    message GenerateTokenReq{
        string accountName=1;
    }

    message GenerateTokenResp{
        bool success=1;
        string token=2;
    }

    service Auth{
        rpc GenerateToken(GenerateTokenReq) returns(GenerateTokenResp);
        rpc ValidateToken(TokenValidateReq) returns(TokenValidateResp);
    }
    ```
2. 进入`/cmd/auth/rpc`目录，执行命令，自动创建源码
    ```cmd
    goctl rpc protoc ./pb/auth.proto --go_out=. --go-grpc_out=. --zrpc_out=.
    ```
3. 修改`cmd/auth/rpc/internal/logic`目录下的`generatetokenlogic.go`,`validatetokenlogic.go`中逻辑，增加`token`生成和`token`验证的逻辑。
4. 运行`cmd/auth/rpc/auth.go`,启动服务端进程。
5. 编写客户端测试程序，`cmd/auth/rpc/test/auth-rpc-client_test.go`，启动向第`4`步启动的服务端发送`rpc`请求。
6. 运行结果如下
    ```log
    === RUN   TestRpcClient
    {"@timestamp":"2023-08-24T12:18:40.994+08:00","caller":"p2c/p2c.go:181","content":"p2c - conn: 127.0.0.1:9000, load: 1850, reqs: 1","level":"stat"}
    resp: &pb.TokenValidateResp{state:impl.MessageState{NoUnkeyedLiterals:pragma.NoUnkeyedLiterals{}, DoNotCompare:pragma.DoNotCompare{}, DoNotCopy:pragma.DoNotCopy{}, atomicMessageInfo:(*impl.MessageInfo)(0xc0000491c8)}, sizeCache:0, unknownFields:[]uint8(nil), Ok:true}
    err: <nil>
    --- PASS: TestRpcClient (0.02s)
    PASS
    ```