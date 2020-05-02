package main

import (
    "fmt"
    "time"
)


func producer(factor int, put chan<- int)  {
    for i:=1; ; i++ {
        put<- i * factor
        time.Sleep(time.Duration(factor) * time.Second)
    }
}

func consumer(get <-chan int)  {
    for val:= range get {
        fmt.Println(val)
    }
}

func main() {
    ch := make(chan int, 100)
    go producer(3, ch)
    go producer(5, ch)
    go consumer(ch)
    consumer(ch)
}
