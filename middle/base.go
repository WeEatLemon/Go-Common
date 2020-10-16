package middle

import (
	resp "github.com/IEatLemons/GoHelper/helper/responses"
	"github.com/go-redis/redis"
)

const (
	KeyLanguage = "language"
	KeyPlatform = "platform"
)

var Middlewares *Middle

type Middle struct {
	Redis         *redis.Client
	InternalHosts string
	InternalIp    string
	Platform      string
	Language      string
}

func Init(R *redis.Client, InternalHost, InternalIp string) {
	Middlewares = &Middle{
		Redis:         R,
		InternalHosts: InternalHost,
		InternalIp:    InternalIp,
	}
	resp.InitResp(Middlewares.Language)
}

func New() *Middle {
	return Middlewares
}
