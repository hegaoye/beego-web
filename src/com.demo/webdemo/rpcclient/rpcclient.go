package rpcclient

import (
	"com.demo/webdemo/rpcprotos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"strconv"
)

/**
发送消息
*/
func ServerToClientMessageResult(toUsername string) bool {
	conn, err := createConnection()
	if err == nil {
		//defer conn.Close()
		client := rpcprotos.NewGoRpcServerInterfaceClient(conn)
		message := rpcprotos.ServerToClientMsgResultParam{
			Code:  110,
			MsgId: 100,
		}
		result, err := client.ServerToClientMessageResult(context.Background(), &message)
		if err != nil || result == nil || result.Code != 200 {
			if err != nil {
				fmt.Println("失败： " + err.Error())
			}
			if result != nil {
				fmt.Println("失败： " + strconv.Itoa(int(result.Code)) + "   " + result.Message)
			}
			return false
		}

	} else {
		fmt.Println("连接RPC服务端失败：" + err.Error())
	}
	return false
}

//连接JavaRPC服务器
func createConnection() (*grpc.ClientConn, error) {
	var err error = nil
	if conn == nil || conn.GetState() != connectivity.Ready {
		address := "127.0.0.1:" + strconv.Itoa(18003)
		conn, err = grpc.Dial(address, grpc.WithInsecure())
	}
	return conn, err
}

//rpc连接
var conn *grpc.ClientConn

//消息模型
type MessageModel struct {
	MsgId      int    `json:"msgId"`      //消息Id
	Type       int8   `json:"type"`       //消息类型
	FromUserId int    `json:"fromUserId"` //发送方UserId
	ToUserId   int    `json:"toUserId"`   //消息接收方UserId
	Content    string `json:"content"`    //消息内容
}
