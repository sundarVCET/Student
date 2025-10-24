package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// Ping Redis to confirm connection
	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println("✅ Connected to Redis!")

}
func ClearAllCache(rdb *redis.Client) {
	ctx := context.Background()
	err := rdb.FlushAll(ctx).Err()
	if err != nil {
		fmt.Println("Error clearing Redis cache:", err)
	} else {
		fmt.Println("✅ All Redis keys deleted")
	}
}
