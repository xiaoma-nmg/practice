package main

import (
    "context"
    "fmt"
    "net"

    "google.golang.org/grpc"

    pb "grpcDemo/hello/proto"
)

type server struct {}

func (s *server)SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplay, error) {
    fmt.Println("from ", in.Name)
    return &pb.HelloReplay{
        Message: "Hello" + in.Name,
    }, nil
}

func (s *server)GetMsg(ctx context.Context, in *pb.HelloRequest) (*pb.HelloMessage, error)  {
    fmt.Println("from ", in.Name)
    return &pb.HelloMessage{
        Msg:"收到了你的消息",
    }, nil
}

func main()  {
    // 监听⽹络
    ln, err := net.Listen("tcp", "127.0.0.1:8888")
    if err != nil {
        fmt.Println("⽹络异常：", err)
        return
    }
    // 创建grpc句柄
    srv := grpc.NewServer()
    // 将server结构体注册到grpc服务中
    pb.RegisterHelloServerServer(srv, &server{})
    // 监听服务
    err = srv.Serve(ln) //阻塞
    if err != nil {
        fmt.Println("监听异常：", err)
        return
    }
}