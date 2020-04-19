package main

import (
    "context"
    "fmt"

    "google.golang.org/grpc"
    pb "grpcDemo/hello/proto"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
    if err != nil {
        fmt.Println("连接服务器失败", err)
    }
    defer conn.Close()
    client := pb.NewHelloServerClient(conn)
    r1, err := client.SayHi(context.TODO(), &pb.HelloRequest{Name:"Wade"})
    if err != nil {
        fmt.Println("can not get SayHi:", err)
    }
    fmt.Println("SayHi 响应：", r1.GetMessage())

    r2, err := client.GetMsg(context.TODO(), &pb.HelloRequest{Name:"James"})
    if err != nil {
        fmt.Println("can not get GetMsg:", err)
    }
    fmt.Println("GetMsg 响应：", r2.GetMsg())
}
