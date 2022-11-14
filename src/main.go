package main

import (
	"GinAPI/src/db"
	"GinAPI/src/middleware"
	"GinAPI/src/sys"
	"github.com/gin-gonic/gin"
)

func main() {
	db.SetupDatabase()

	r := gin.Default()

	r.Use(middleware.CustomParamBuilder())
	sys.SetupRouteTable(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
