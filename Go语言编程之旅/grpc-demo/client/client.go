package main

import (
	"context"
	"google.golang.org/grpc"
	pb "github.com/wsloong/grpc-demo/proto"
	"io"
	"log"
)


func main() {
	conn, _ := grpc.Dial("127.0.0.1:8001", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// Unary Rpc: 一元 RPC的客户端
	if err := SayHello(client); err != nil {
		log.Fatalf("SayHello err: %v", err)
	}

	// Server-side streaming RPC: 服务端流式 RPC
	//if err := SayList(client, &pb.HelloRequest{Name: "wsl"}); err != nil {
	//	log.Fatalf("SayList err: %v", err)
	//}

	//// Client-side streaming RPC: 客户端端流式 RPC
	//if err := SayRecord(client, &pb.HelloRequest{Name: "wsl"}); err != nil {
	//	log.Fatalf("SayRecord err: %v", err)
	//}

	// Bidirectional streaming RPC: 双向流式 RPC
	//if err := SayRoute(client, &pb.HelloRequest{Name:"wsl"}); err != nil {
	//	log.Fatalf("SayRoute err: %v", err)
	//}
}

func SayHello(client pb.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "wsl"})
	if err != nil {
		return err
	}

	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayList(context.Background(), r)
	if err != nil {
		return nil
	}

	for {
		// Recv() 封装了 ClientStream.RecvMsg； RecvMsg 从流中读取完整的 gRPC消息体
		//  RecvMsg 是阻塞等待的
		// 当流成功或者结束时候，RecvMsg 会返回 io.EOF
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: %v", resp)
	}
	return nil
}

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayRecord(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		if err := stream.Send(r); err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("resp err: %v", resp)
	return nil
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayRoute(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		if err := stream.Send(r); err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp err: %v", resp)
	}

	if err := stream.CloseSend(); err != nil {
		return err
	}
	return nil
}