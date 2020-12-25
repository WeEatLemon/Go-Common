package middle

import (
	"github.com/IEatLemons/GoHelper/language"
	"github.com/gin-gonic/gin"
)

func (M *Middle) AuthLanguage() func(c *gin.Context) {
	return func(c *gin.Context) {
		Language := c.GetHeader(KeyLanguage)
		if Language != language.EnLan && Language != language.ZhCnLan {
			Language = language.ZhCnLan
		}
		c.Set(KeyLanguage, Language)
		M.Language = Language
	}
}
