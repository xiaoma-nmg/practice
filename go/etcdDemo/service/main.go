package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"etcdDemo/etcdconnect"
	pb "etcdDemo/proto"
	"google.golang.org/grpc"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
)

const host = "127.0.0.1"

var (
	ServiceName = flag.String("ServiceName", "hello_service", "service name")
	Port        = flag.Int("Port", 3000, "listening port")
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address")
)

type server struct{}

func (s *server) SayHi(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("[SayHi] recive [name:%s, age:%d]\n", in.Name, in.Age)
	return &pb.HelloResponse{
		Response: fmt.Sprintf("Hi %s, you age is %d", in.Name, in.Age),
	}, nil
}
func (s *server) GetMessage(ctx context.Context, in *pb.SendMessage) (*pb.RecvMessage, error) {
	log.Printf("[GetMessage] recive message [%s]\n", in.Message)
	return &pb.RecvMessage{
		Message: "this is for your",
	}, nil
}

func main() {
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", host, *Port)
	// 监听网络
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	// 注册server
	srv := grpc.NewServer()
	defer srv.GracefulStop()
	pb.RegisterHelloServer(srv, &server{})

	// 服务器注册到etcd服务器
	go etcdconnect.Register(*EtcdAddr, *ServiceName, addr, 5)

	// 接收到系统退出server的信号时，将server从etcd注销
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		sign := <-ch
		etcdconnect.UnRegister(*ServiceName, addr)
		if i, ok := sign.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	// 启动server
	_ = srv.Serve(listen)
}
