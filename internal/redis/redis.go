package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "",
		DB:       0,
	})

	// проверка коннекта
	if err := Rdb.Ping(Ctx).Err(); err != nil {
		panic(fmt.Sprintf("Redis connection failed: %v", err))
	}
	fmt.Println("✅ Connected to Redis")
}
