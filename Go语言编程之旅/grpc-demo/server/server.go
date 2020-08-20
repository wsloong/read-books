package main

import (
	"context"
	pb "github.com/wsloong/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)
type GreeterServer struct {}

// Unary RPC: 一元RPC(单次RPC),客户端发起一次普通的 RPC 请求
// 查看 .proto 文件  SayHello定义
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello world"}, nil
}

// Server-side streaming RPC: 服务端流式 RPC
// 服务端流式RPC是一个单向流，值Server为Stream，客户端为普通的一元 RPC 请求
// 客户端发起一次普通的RPC请求，服务器通过流式响应多次发送数据集
//  查看 .proto 文件  SayList 定义
func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	var err error
	for n := 0; n <= 6; n++ {
		err = stream.Send(&pb.HelloReply{Message: "hello list"})
	}
	return err
}

// Client-side streaming RPC: 客户端流式RPC
// 客户端 通过流式发起多次 RPC 请求给服务器，服务器仅发起一次响应给客户端
//  查看 .proto 文件  SayRecord 定义
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		// 当流关闭时候，需要通过 stream.SendAndClose 将最终的响应结果发送给客户端
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "say.record"})
		}
		if err != nil {
			return err
		}
		log.Printf("resp: %v", resp)
	}
	return nil
}

// Bidirectional streaming RPC: 双向流式RPC
// 客户端以流式的方式发起请求，服务器同样以流式的方式响应请求
// 首个请求一定是客户端发起的
func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		if err := stream.Send(&pb.HelloReply{Message: "say.route"});  err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("resp: %v", resp)

	}
}

func main() {
	// 实例化对象
	server := grpc.NewServer()
	// 注册到gRPC Server的内部注册中心
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8001")
	server.Serve(lis)
}