package main

import (
    "context"
    "fmt"
    "time"

    "github.com/micro/go-micro/v2"
    pb "micro/protoBased/proto"
)

func main() {
    service := micro.NewService()
    service.Init()

    // 创建专属于Greeter服务的配套客户端
    greeter:= pb.NewGreeterService("greeter.service", service.Client())

    ticker := time.NewTicker(3 * time.Second)

    for {
        response, err := greeter.Hello(context.Background(), &pb.HelloRequest{Name: "wade"})
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(response.Greeting)

        response2, err := greeter.Hello2(context.Background(), &pb.HelloRequest{Name: "Wade"})
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Println(response2.Greeting)

        <-ticker.C
    }
}
