package middle

import (
	resp "github.com/WeEatLemon/Go-Common/helper/responses"
	"github.com/gin-gonic/gin"
)

const (
	Default = "Test"
)

func (M *Middle) AuthPlatform() func(c *gin.Context) {
	return func(c *gin.Context) {
		platform := c.GetHeader(KeyPlatform)
		if platform == "" {
			resp.NewResp().ParamErrRep(c, &resp.Elem{
				Code: resp.ParameterError,
				Msg:  "Request header must bring platform",
			})
			c.Abort()
			return
		}
		switch platform {
		case Default:
		default:
			resp.NewResp().ParamErrRep(c, &resp.Elem{
				Code: resp.ParameterError,
				Msg:  "Unrecognized platform",
			})
			c.Abort()
			return
		}
		c.Set(KeyPlatform, platform)
		M.Platform = platform
	}
}
