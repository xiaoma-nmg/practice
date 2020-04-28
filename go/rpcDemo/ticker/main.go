package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func main() {
    // 心跳发生器
    ticker := time.NewTicker(time.Second * 2)
    go func() {
        for {
            fmt.Println("heart beat.")
            <-ticker.C
        }
    }()
    scan := bufio.NewScanner(os.Stdin)
    for scan.Scan() {
        fmt.Println(scan.Text())
    }
}
