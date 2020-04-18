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

func main()  {
    if len(os.Args) != 2 {
        fmt.Println("usage:", os.Args[0], "server:port")
        os.Exit(1)
    }
    serverAddress := os.Args[1]
    client, err := rpc.Dial("tcp", serverAddress)
    if err != nil {
        log.Fatal("dialing:", err)
    }

    args := Args{
        A: 333,
        B: 78,
    }

    var reply int
    err = client.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("airth error:", err)
    }
    fmt.Printf("%d*%d=%d\n", args.A, args.B, reply)

    var quo Quotient
    err = client.Call("Arith.Divide", args, &quo)
    if err != nil {
        log.Fatal("airth error:", err)
    }
    fmt.Printf("%d/%d=%d rem %d\n", args.A, args.B, quo.Quo, quo.Rem)
}
