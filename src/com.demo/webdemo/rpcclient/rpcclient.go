package rpcclient

import (
	"com.demo/webdemo/rpcprotos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
)

// ServerToClientMessageResult 發送消息
func ServerToClientMessageResult(toUsername string) bool {
	conn, err := createConnection()
	if err == nil {
		//defer conn.Close()
		client := rpcprotos.NewGoRpcServerInterfaceClient(conn)
		message := rpcprotos.ServerToClientMsgResultParam{
			Code:  110,
			MsgId: 100,
		}
		result, _ := client.ServerToClientMessageResult(context.Background(), &message)
		fmt.Printf("服務器返回結果：", result)

	} else {
		fmt.Println("连接RPC服务端失败：" + err.Error())
	}
	return false
}

func createConnection() (*grpc.ClientConn, error) {
	var err error = nil
	if conn == nil || conn.GetState() != connectivity.Ready {
		conn, err = grpc.Dial("127.0.0.1:18003", grpc.WithInsecure())
	}
	return conn, err
}

//rpc连接
var conn *grpc.ClientConn

// MessageModel 消息模型
type MessageModel struct {
	MsgId      int    `json:"msgId"`      //消息Id
	Type       int8   `json:"type"`       //消息类型
	FromUserId int    `json:"fromUserId"` //发送方UserId
	ToUserId   int    `json:"toUserId"`   //消息接收方UserId
	Content    string `json:"content"`    //消息内容
}
