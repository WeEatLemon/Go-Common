package middle

import (
	"encoding/json"
	"fmt"
	resp "github.com/IEatLemons/GoHelper/helper/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

/* 通用访问限制 */
func (M *Middle) Cors() func(c *gin.Context) {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, lan, token, platform")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

/* 访问限制 */
func (M *Middle) InternalCors() func(c *gin.Context) {
	return func(c *gin.Context) {
		var hosts []string
		err := json.Unmarshal([]byte(M.InternalHosts), &hosts)
		if err != nil {
			resp.NewResp().ServerErrRep(c, &resp.Elem{
				Code: resp.UnknownError,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}

		vhost := "localhost"
		isOrigin := false
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		fmt.Println("Cli Host:", origin)
		for _, host := range hosts {
			if host == "*" || host == origin {
				isOrigin = true
				vhost = host
				goto SetOrigin
			}
		}

		if !isOrigin {
			resp.NewResp().ParamErrRep(c, &resp.Elem{
				Code: resp.UntrustedSource,
			})
			c.Abort()
			return
		}

	SetOrigin:
		c.Header("Access-Control-Allow-Origin", vhost)
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, lan, platform")
		c.Header("Access-Control-Allow-Methods", "POST, PUT, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
