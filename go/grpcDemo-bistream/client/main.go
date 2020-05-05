package main

import (
    "bufio"
    "context"
    "io"
    "log"
    "os"

    "google.golang.org/grpc"
    pb "grpcDeom-bistream/proto"
)

func main() {
    conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
    if err != nil {
        log.Printf("连接失败：[%v]", err)
        return
    }
    defer conn.Close()
    // 声明客户端
    client := pb.NewChatClient(conn)
    // 声明 context
    ctx := context.Background()
    // 创建双向数据流
    stream, err := client.BidStream(ctx)
    if err != nil {
        log.Printf("创建数据流失败：[%v]\n", err)
        return
    }

    // 用一个goroutine 接收命令行的输入
    go func() {
        log.Println("请输入消息...")
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            if err := stream.Send(&pb.Request{Input:scanner.Text()}); err != nil {
                return
            }
        }
    }()

    for {
        // 接收从服务器返回的数据流并打印
        reply, err := stream.Recv()
        if err == io.EOF {
            log.Println("收到服务端的结束信号")
            break
        }
        if err != nil {
            log.Printf("接收出错：[%v]\n", err)
            break
        }
        log.Printf("收到消息：[%v]\n", reply.Output)
    }
}
