package pkpk

import (
	"github.com/gin-gonic/gin"
)

func GetRequestParam(c *gin.Context) []string {
	val, exist := c.Get("arr")
	arr := val.([]string)
	if !exist {
		panic("Error")
	}

	return arr
}
