
博客网站24小时文章排行榜实现：

需求：
    1。用redis的zset实现
    2。24小时的
    3。排行  

数据结构设计
    zset:
        sorted sort:有序的集合 ---> key:(a 90, b 80, c 70 ...)
        
    KEY ：         当天（24小时）的文章点击数标识   (分段式的key)
    集合中的元素：   每个文章的ID
    元素的分数：     每个文章被点击的次数
业务设计：
    1。点击文章，阅读数加1  
    
    zincrby key 1 member
      
    2。文章排行

    zrevrange key start end withscores

