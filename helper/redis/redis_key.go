package redis

const (
	UserToken     = "LoginToken:"    // 用户-token
	Transaction   = "Tx:"            // 处理交易 Token
	CachePayOrder = "CachePayOrder:" // 处理交易 Token
)

/* 获取支付缓存 Key */
func GetCachePayOrderKey(OrderNo string) string {
	return CachePayOrder + OrderNo
}

func GetTransactionKey(OrderNo string) string {
	return Transaction + OrderNo
}

func GetLoginTokenKey(Token string) string {
	return UserToken + Token
}
