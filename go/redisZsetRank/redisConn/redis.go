package redisConn

import (
    "fmt"

    "github.com/go-redis/redis"
)

var (
    Client *redis.Client
)

const (
    HOST = "127.0.0.1"
    PORT = 6379
)

// 初始化连接
func InitRedis() (err error) {
    Client = redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%d", HOST, PORT),
    })
    
    _, err = Client.Ping().Result()
    if err != nil {
        return nil
    }
    return nil
}
