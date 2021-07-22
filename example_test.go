package redis_rate_test

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

func ExampleNewLimiter() {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
	})

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })
	_ = rdb.FlushDB(ctx).Err()

	limiter := redis_rate.NewLimiter(rdb)

	for i := 0; i < 20; i++ {
		res, err := limiter.AllowN(ctx, "project:123", redis_rate.Limit{
			Rate:   2,
			Period: time.Second,
			Burst:  40.0,
		}, 1.5)
		if err != nil {
			panic(err)
		}
		fmt.Println("allowed", res.Allowed, "remaining", res.Remaining, "retry-after", res.RetryAfter)
		if i == 10 {
			fmt.Println("time.Sleep(1 * time.Second)")
			time.Sleep(1 * time.Second)
		}
	}

	// Output: allowed 1 remaining 9
}

// insert into third_party_searches values(1, "2021-07-16 07:35:00", "2021-07-16 07:35:00", "123", 1, "reference_key", "endpoint")
// insert into stores values(1, 1, "2021-07-16 07:35:00", "2021-07-16 07:35:00");
// insert into search_enabled_tracks values(1, 1, "123", "2021-07-16 07:35:00", "2021-07-16 07:35:00");
