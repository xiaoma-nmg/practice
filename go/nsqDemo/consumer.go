package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

type MyHandler struct {
	Title string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

func initConsumer(topic string, channel string, addr string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second

	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, error:%v\n", err)
		return err
	}

	consumer := &MyHandler{Title: "mqTest"}
	c.AddHandler(consumer)

	// 通过nsqd查询
	//if err := c.ConnectToNSQD(addr); err != nil
	// 通过lookupd查询
	if err := c.ConnectToNSQLookupd(addr); err != nil {
		return err
	}
	return nil
}

func main() {
	err := initConsumer("nsq_demo_topic", "first", "127.0.0.1:4161")
	if err != nil {
		fmt.Printf("init consumer failed, error:%v\n", err)
		return
	}

	sign := make(chan os.Signal)        // 定义一个信号的通道
	signal.Notify(sign, syscall.SIGINT) // 转发键盘中断信号到sign
	<-sign                              // 阻塞
}
