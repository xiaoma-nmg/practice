package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/nsqio/go-nsq"
)

const nsqServer  = "127.0.0.1:4150"

var producer *nsq.Producer

func initProducer(addr string) (err error) {
    config := nsq.NewConfig()
    producer, err = nsq.NewProducer(addr, config)
    if err != nil {
        fmt.Printf("Create producer failed, err: %v\n", err)
        return err
    }
    return nil
}

func main() {
    err := initProducer(nsqServer)
    if err != nil {
        fmt.Printf("init Producer error: %v\n", err)
        return
    }

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("please input something:")
    for scanner.Scan() {
        data := scanner.Text()
        data = strings.TrimSpace(data)

        //输入Q 退出
        if strings.ToUpper(data) == "Q" {
            fmt.Println("exit program")
            break
        }

       if  err := producer.Publish("nsq_demo_topic", []byte(data)); err != nil {
           fmt.Printf("publish message to nsq failed, err :%v\n", err)
           continue
       }
    }
}
