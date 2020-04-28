package main

import (
    "context"
    "fmt"
    "net"

    "google.golang.org/grpc"
    pb "grpcDemo/proto"
)

type server struct {}



func(s *server)SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
    return &pb.HelloReplay{
        Message: "Hello " + in.Name,
    }, nil
}
func(s *server)GetMsg(ctx context.Context, in *pb.HelloRequest) (*pb.HelloMessage, error) {
    return &pb.HelloMessage{
        Msg: "you have my word!",
    }, nil
}

func main() {
    ln, err := net.Listen("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println("create socket error: ", err)
        return
    }

    // 创建grpc句柄
    srv := grpc.NewServer()
    // 将server结构体注册到grpc服务中
    pb.RegisterHelloServerServer(srv, &server{})
    // 监听服务
    if err := srv.Serve(ln); err != nil {
        fmt.Println("server listen error: ", err)
        return
    }
}
