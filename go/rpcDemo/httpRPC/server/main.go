package main

import (
    "errors"
    "fmt"
    "net/http"
    "net/rpc"
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
    rpc.HandleHTTP()

    if err := http.ListenAndServe("127.0.0.1:9999", nil); err != nil {
        fmt.Println(err)
    }
}
