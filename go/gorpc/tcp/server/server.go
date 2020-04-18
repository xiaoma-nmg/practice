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

func (t *Arith)Multiply(args *Args, reply *int) error  {
    *reply = args.A * args.B
    return nil
}

func (t *Arith)Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

func checkError(err error)  {
    if err != nil {
        fmt.Println("Fatal error", err.Error())
    }
}

func main()  {
    ar := new(Arith)
    rpc.Register(ar)

    tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)

    for {
        fmt.Printf("accept at %v\n", time.Now())
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        fmt.Printf("get at %v\n", time.Now())
        go rpc.ServeConn(conn)
    }
}