package middle

import (
	"encoding/json"
	"fmt"
	resp "github.com/WeEatLemon/Go-Common/helper/responses"
	"github.com/gin-gonic/gin"
)

func (M *Middle) AuthIP() func(c *gin.Context) {
	return func(c *gin.Context) {
		CliIP := c.ClientIP()
		fmt.Println("CliIP:", CliIP)
		var Ips []string
		err := json.Unmarshal([]byte(M.InternalIp), &Ips)
		if err != nil {
			resp.NewResp().ServerErrRep(c, &resp.Elem{
				Code: resp.UnknownError,
				Msg:  err.Error(),
			})
			c.Abort()
			return
		}

		isGoOut := false
		for _, Ip := range Ips {
			if Ip == "*" || Ip == CliIP {
				isGoOut = true
				goto Next
			}
		}
		if !isGoOut {
			resp.NewResp().ParamErrRep(c, &resp.Elem{
				Code: resp.UntrustedSource,
			})
			c.Abort()
			return
		}

	Next:
		c.Next()
	}
}
