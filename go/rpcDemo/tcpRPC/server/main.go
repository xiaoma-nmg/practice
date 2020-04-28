package main

import (
    "errors"
    "fmt"
    "net"
    "net/rpc"
    "time"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (a *Arith)Multiply(args *Args, reply *int) error  {
    fmt.Println("Multiply receive a rpc Call args is ", args)
    *reply = args.A * args.B
    return nil
}

func (a *Arith)Divide(args *Args, quo *Quotient) error  {
    fmt.Println("Divide receive a rpc Call args is ", args)
    if args.B == 0{
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

func main() {
    arith := new(Arith)
    _ = rpc.Register(arith)

    tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
    if err != nil {
        fmt.Println("create tcp socket error: ", err)
        return
    }
    listen , err := net.ListenTCP("tcp", tcpAddr)
    if err != nil {
        fmt.Println("listen tcp socket error: ", err)
        return
    }

    for {
        fmt.Printf("accept at %v\n", time.Now())
        conn, err := listen.Accept()
        if err != nil {
            continue
        }
        fmt.Println("receive from ", conn.RemoteAddr())
        go rpc.ServeConn(conn)
    }
}
