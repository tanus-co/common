//grpc 	请求封装
package rpc

import (
	"google.golang.org/grpc"
)

//grpc 请求
func RpcRequest(address string, method func(conn *grpc.ClientConn) (interface{}, error)) (interface{}, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return method(conn)
}
