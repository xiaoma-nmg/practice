package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"etcdDemo/etcdconnect"
	pb "etcdDemo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	ServiceName = flag.String("ServiceName", "hello_service", "service name")
	EtcdAddr    = flag.String("EtcdAddr", "127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()

	r := etcdconnect.NewetcdResolver(*EtcdAddr)
	resolver.Register(r)

	//conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	// 客户端连接服务器
	// 用字符串服务发现连接服务器
	conn, err := grpc.Dial(r.Scheme()+"://author/"+*ServiceName,
		grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 获取grpc句柄
	client := pb.NewHelloClient(conn)

	ticker := time.NewTicker(2 * time.Second)

	for {
		r1, err := client.SayHi(context.Background(), &pb.HelloRequest{
			Name: "Tom",
			Age:  23,
		})
		if err != nil {
			log.Printf("[SayHi] error:%#v\n", err)
			return
		}
		fmt.Printf("SayHi result is [%s] \n", r1.GetResponse())

		r2, err := client.GetMessage(context.Background(), &pb.SendMessage{
			Message: "this is a test",
		})
		if err != nil {
			log.Printf("[GetMessage] error:%#v\n", err)
			return
		}
		fmt.Printf("GetMessage result is [%s]\n", r2.GetMessage())

		<-ticker.C
	}

}
