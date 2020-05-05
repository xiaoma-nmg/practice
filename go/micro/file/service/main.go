package main

import (
    "context"
    "fmt"
    "io/ioutil"

    "github.com/micro/go-micro/v2"
    "github.com/micro/go-micro/v2/errors"
    pb "micro/file/proto"
)

type File struct {}

func (f *File) File(ctx context.Context, file pb.File_FileStream) error {
    // 将收到的文件存储在临时文件夹中
    // File_FileStream类型有Recv函数，源源不断从客户端收到slice的消息
    temp, err := ioutil.TempFile("", "micro")
    if err != nil {
        return errors.InternalServerError("file.service", err.Error())
    }

    for {
        b, err := file.Recv() // b 被自动unmarshall成FileSlice类型
        if err != nil {
            return errors.InternalServerError("file.service", err.Error())
        }
        // Len 是-1 表明发送完毕
        if b.Len == -1 {
            break
        }

        if _, err := temp.Write(b.Byte); err != nil { // 流式的append到文件中
            return errors.InternalServerError("file.service", err.Error())
        }
    }

    fmt.Println(temp.Name()) //生成的临时文件名
    return file.SendMsg(&pb.FileSliceMsg{
        FileName:temp.Name(),
    })
}

func main() {
    service := micro.NewService(
        micro.Name("file.service"),
    )
    service.Init()
    pb.RegisterFileHandler(service.Server(), &File{})

    service.Run()
}
