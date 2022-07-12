package rpcserver

import (
	"com.demo/webdemo/rpcprotos"
	"context"
	"fmt"
)

type GoRpcServerInterfaceServerImpl struct {
	rpcprotos.UnimplementedGoRpcServerInterfaceServer
}

func (*GoRpcServerInterfaceServerImpl) IMLogin(ctx context.Context, req *rpcprotos.LoginParam) (*rpcprotos.GoBaseResult, error) {

	account := req.Account
	fmt.Println("IM登录，用户名：" + account)
	return &rpcprotos.GoBaseResult{
		Code:    200,
		Message: "登陆成功",
	}, nil

}

func (*GoRpcServerInterfaceServerImpl) ServerToClientMessageResult(ctx context.Context,
	req *rpcprotos.ServerToClientMsgResultParam) (*rpcprotos.GoBaseResult, error) {
	fmt.Printf("收到來自客戶端消息 ： ", req.String())
	return &rpcprotos.GoBaseResult{
		Code:    200,
		Message: "OK",
	}, nil
}
