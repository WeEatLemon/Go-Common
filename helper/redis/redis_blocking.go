package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"time"
)

func BlockingKey(Redis *redis.Client, Key string, expiration time.Duration) (err error) {
	val := Redis.Get(Key).Val()
	if val != "" {
		return errors.New("")
	}
	Redis.Set(Key, "Blocking Key", expiration)
	return
}
