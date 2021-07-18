package main

import (
	"fmt"
	"runtime"

	"num_calculator/ctrl"
	"num_calculator/middleware/cors"
	"num_calculator/middleware/token"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	port := *pflag.StringP("port", "p", "8080", "port for listen")
	pflag.Parse()

	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	g.Use(cors.Cors())
	g.Use(token.TokenVerify())
	g.POST("/cal", ctrl.Calculate)
	g.GET("/dl", ctrl.Download)

	g.Run(fmt.Sprintf(":%s", port))
}
