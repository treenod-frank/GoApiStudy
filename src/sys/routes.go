package sys

import (
	"GinAPI/src/controllers/todoCtrl"
	"github.com/gin-gonic/gin"
)

type RouteInfo struct {
	Path    string
	Handler func(ctx *gin.Context)
}

func SetupRouteTable(r *gin.Engine) {
	routePostInfos := []RouteInfo{
		{
			"/getTodos", todoCtrl.GetTodos,
		},
		{
			"/addTodos", todoCtrl.AddTodos,
		},
		{
			"/updateTodos", todoCtrl.UpdateTodos,
		},
	}

	for _, routeInfo := range routePostInfos {
		r.POST("/pokopoko"+routeInfo.Path+".php", routeInfo.Handler)
	}
}
