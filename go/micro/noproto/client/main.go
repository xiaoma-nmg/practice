package main

import (
    "context"
    "fmt"

    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/client"
)

func main() {
    service := micro.NewService()
    service.Init()
    c := service.Client()

    // 客户端的请求过程：
    // 通过指定service的名称进行服务发现
    // 通过服务发现选取到服务中的某个node
    // 给node发送Grpc请求
    // 得到响应
    request := c.NewRequest("service.greeter", "Greeter.Hello", "Wade",
        client.WithContentType("application/json"))

    var response string
    if err := c.Call(context.Background(), request, &response); err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(response)
    request2 := c.NewRequest("service.greeter", "Greeter.Response", "test",
        client.WithContentType("application/json"))
    if err := c.Call(context.Background(), request2, &response); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(response)
}
