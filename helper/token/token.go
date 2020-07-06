package token

import (
	"github.com/go-redis/redis"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n uint) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateLoginToken(Redis *redis.Client, Data interface{}) (token string) {
	token = RandStringRunes(32)

	_, err := Redis.Get(token).Result()
	if err != redis.Nil {
		return CreateLoginToken(Redis, Data)
	}

	Redis.Set(token, Data, time.Second)
	return
}
