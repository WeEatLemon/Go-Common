package redis

const (
	UserToken      = "login_token_"     // 用户-token
	GoogleKey      = "GG:KEY:"          // 用户-token
	Transaction    = "Tx:"              // 处理交易 Token
	CachePayOrder  = "CachePayOrder:"   // 处理交易 Token
	OrderCountDown = "ping:order:down:" //拼团倒计时
)

/* 获取支付缓存 Key */
func GetCachePayOrderKey(OrderNo string) string {
	return CachePayOrder + OrderNo
}
