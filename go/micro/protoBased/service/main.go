package main

import (
    "context"
    "fmt"

    "github.com/micro/go-micro/v2"
    pb "micro/protoBased/proto"
)

type Greeter struct {}

func (g *Greeter)Hello(ctx context.Context, in *pb.HelloRequest, out *pb.HelloResponse) error {
    fmt.Println("Hello function")
    out.Greeting = "你好， " + in.Name
    return nil
}

func (g *Greeter)Hello2(ctx context.Context, in *pb.HelloRequest, out *pb.HelloResponse) error {
    fmt.Println("this is Hello2 function")
    out.Greeting = "Hello2 收到了你的消息。"
    return nil
}

func main() {
    service := micro.NewService(
        micro.Name("greeter.service"),
        micro.Version("v1"),
    )
    service.Init()

    pb.RegisterGreeterHandler(service.Server(), &Greeter{})

    service.Run()
}
