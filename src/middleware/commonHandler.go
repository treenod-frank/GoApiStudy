package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func CustomParamBuilder() gin.HandlerFunc {
	return func(c *gin.Context) { // 미들웨어 등록
		q, has := c.GetPostForm("q")
		if !has {
			panic("Error")
		}
		arr := strings.Split(string(q), ",")
		c.Set("arr", arr)
	}
}
