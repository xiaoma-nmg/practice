package main

import (
    "fmt"
    "log"
    "net/rpc"
    "os"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: ", os.Args[0], "server")
        os.Exit(1)
    }
    serverAddress := os.Args[1]

    client, err := rpc.Dial("tcp", serverAddress)
    if err != nil {
        log.Fatal("dialing server error: ", err)
    }

    args := Args{17, 8}
    var reply int
    if err := client.Call("Arith.Multiply", args, &reply); err != nil {
        log.Fatal("call Arith.Multiply error: ", err)
    }
    fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)

    var quo Quotient
    if err := client.Call("Arith.Divide", args, &quo); err != nil {
        log.Fatal("call Arith.Divide error: ", err)
    }
    fmt.Printf("Arith: %d / %d = %d remainder %d\n", args.A, args.B, quo.Quo, quo.Rem);
}
