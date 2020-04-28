package main

import (
    "io"
    "log"
    "net"
    "strconv"

    "google.golang.org/grpc"
    pb "grpcDeom-bistream/proto"
)

// Streamer 服务端
type Streamer struct{}

// BidStream 实现了 ChatServer 接口中定义的 BidStream 方法
func (s *Streamer) BidStream(stream pb.Chat_BidStreamServer) error {
    ctx := stream.Context()
    for {
        select {
        case <-ctx.Done():
            log.Println("收到客户端通过context发出的终止信号")
            return ctx.Err()
        default:
            // 接收客户端发来的消息
            input, err := stream.Recv()
            if err == io.EOF {
                log.Println("客户端发送的数据流结束")
                return nil
            }
            if err != nil {
                log.Println("接收数据出错： ", err)
                return err
            }

            switch input.Input {
            case "end":
                log.Println("收到结束指令")
                if err := stream.Send(&pb.Response{Output:"收到结束指令"}); err != nil {
                    return err
                }
                return nil
            case "back stream":
                log.Println("收到返回数据流指令")
                for i:=0; i<10; i++ {
                    if err := stream.Send(&pb.Response{Output:"数据流 #" + strconv.Itoa(i)}); err != nil {
                        return err
                    }
                }
            default:
                log.Printf("[收到消息]：%s\n", input.Input)
                if err := stream.Send(&pb.Response{Output:"服务端返回： " + input.Input}); err != nil {
                    return err
                }
            }
        }
    }
}


func main()  {
    log.Println("启动服务端...")
    server := grpc.NewServer()
    // 注册 ChatServer
    pb.RegisterChatServer(server, &Streamer{})
    address, err := net.Listen("tcp", ":9999")
    if err != nil {
        panic(err)
    }
    if err := server.Serve(address); err != nil {
        panic(err)
    }
}
