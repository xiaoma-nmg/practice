package main

import (
    "context"
    "fmt"

    "github.com/micro/go-micro/v2"
)

type Greeter struct {

}

func (g *Greeter) Hello (c context.Context, name *string, msg *string) error {
    fmt.Println("this is Hello function")
    *msg = "hi " + *name + ", this is response for you"
    return nil
}

func main()  {
    service := micro.NewService(micro.Name("service.My.First.Test"))
    service.Init()

    // RegisterHandler 用作 把某个service的Server 和 处理的类型进行 bind
    _ = micro.RegisterHandler(service.Server(), new(Greeter))
    _ = service.Run()
}