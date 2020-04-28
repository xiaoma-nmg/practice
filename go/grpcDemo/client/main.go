package main

import (
    "context"
    "fmt"

    "google.golang.org/grpc"
    pb "grpcDemo/proto"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
    if err != nil {
        fmt.Println("connection server error: ", err)
        return
    }
    defer conn.Close()

    // 获取grpc句柄
    c := pb.NewHelloServerClient(conn)
    reply, err := c.SayHi(context.TODO(), &pb.HelloRequest{Name:"Tom"})
    if err != nil {
        fmt.Println("SayHi call error: ", err)
        return
    }
    fmt.Printf("call SayHi reply is [%s]\n", reply)

    mess, err := c.GetMsg(context.TODO(), &pb.HelloRequest{})
    if err != nil {
        fmt.Println("GetMsg call error: ", err)
        return
    }
    fmt.Printf("call GetMsg result is [%s]\n", mess)
}
