# Rate limiting for go-redis

"github.com/go-redis/redis/v8"
"github.com/go-redis/redis_rate/v9"
的两个包不支持向桶填充浮点型，现在改成可以支持浮点型