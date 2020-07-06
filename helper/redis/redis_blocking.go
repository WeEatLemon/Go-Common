package redis

import (
	"errors"
	resp "github.com/IEatLemons/GoBase/common/helper/responses"
	"github.com/go-redis/redis"
	"time"
)

func BlockingKey(Redis *redis.Client, Key string, expiration time.Duration, Resp *resp.Resp) (err error) {
	val := Redis.Get(Key).Val()
	if val != "" {
		return errors.New(Resp.GetMsgStr(resp.FrequencyTooFast))
	}
	Redis.Set(Key, "Blocking Key", expiration)
	return
}
