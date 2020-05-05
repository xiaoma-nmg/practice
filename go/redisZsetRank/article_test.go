package redisZsetRank

import (
    "testing"

    "redisZsetRank/logic"
    "redisZsetRank/redisConn"
)

func TestRedisZset(t *testing.T)  {
    redisConn.InitRedis()
    logic.IncArticleReadCount("1")
    logic.ArticleTopN(int64(3))
    logic.IncArticleReadCount("1")
    logic.ArticleTopN(int64(3))
    logic.IncArticleReadCount("2")
    logic.IncArticleReadCount("3")
    logic.ArticleTopN(int64(3))
    logic.IncArticleReadCount("1")
    logic.IncArticleReadCount("2")
    logic.IncArticleReadCount("2")
    logic.ArticleTopN(int64(3))
    logic.IncArticleReadCount("3")
    logic.IncArticleReadCount("3")
    logic.IncArticleReadCount("3")
    logic.IncArticleReadCount("3")
    logic.IncArticleReadCount("3")
    logic.IncArticleReadCount("3")
    logic.IncArticleReadCount("3")
    logic.IncArticleReadCount("3")
    logic.ArticleTopN(int64(3))
}
