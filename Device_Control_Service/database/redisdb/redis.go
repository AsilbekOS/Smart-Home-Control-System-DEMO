package redis

import (
	"context"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis bilan ulanishda xato: %v", err)
	}

	log.Println("Redis serveriga muvaffaqiyatli keshlandi:", pong)
	return rdb, nil
}
