package main

import (
    "context"
    "fmt"

    "github.com/micro/go-micro/v2"
)

type Greeter struct {}

func (g *Greeter)Hello(ctx context.Context, name *string, msg *string) error  {
    fmt.Println("A new request comes")
    *msg = "你好， " + *name + " 这是一条response"
    return nil
}

func (g *Greeter)Response(ctx context.Context, arg *string, msg *string) error  {
    fmt.Println("this is Response function")
    *msg = "收到您的消息"
    return nil
}

func main() {
    service := micro.NewService(
        micro.Name("service.greeter"),
    )

    service.Init()
    _ = micro.RegisterHandler(service.Server(), &Greeter{})
    _ = service.Run()
}
