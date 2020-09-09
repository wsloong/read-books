package main

import (
	"context"
	pb "github.com/wsloong/tag-service/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	ctx := context.Background()

	// grpc.WithBlock 设置在 DialContext 方法时立即和服务器链接
	clientConn, err := GetClientConn(ctx, "localhost:8009", []grpc.DialOption{grpc.WithBlock()})
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	defer clientConn.Close()

	// 初始化指定 RPC Proto Service 的客户端实例对象
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	// 发起指定 RPC 方法的调用
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Fatalf("tagServiceClient.GetTagList err: %v", err)
	}
	log.Printf("resp: %v", resp)
}

// 创建给定目标的客户端链接。非加密模式，所以要调用 grpc.WithInsecure 方法来禁用此 ClientConn 的传输安全性验证
func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}