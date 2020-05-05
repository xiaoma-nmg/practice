package logic

import (
    "fmt"
    "time"

    "redisZsetRank/redisConn"
)

const (
    KeyArticleCount = "blog:article:count:%s" // 24小时文章阅读数 eq: blog:article:count:20200505
)


// 点击文章，阅读数加1
// 执行 zincrby key 1 member
func IncArticleReadCount(articleID string) error {
    todayStr := time.Now().Format("20060102")
    key := fmt.Sprintf(KeyArticleCount, todayStr)
    return redisConn.Client.ZIncrBy(key, 1, articleID).Err()
}


// 打印排名最高的前N篇文章
func ArticleTopN(n int64)  {
    // 1.  zrevrange key 0 n-1
    todayStr := time.Now().Format("20060102")
    key := fmt.Sprintf(KeyArticleCount, todayStr)
    idStrs, err := redisConn.Client.ZRevRange(key, 0, n-1).Result()
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, ID := range idStrs {
        fmt.Println(ID)
    }
    fmt.Println("-------------------------------------")
    return
}