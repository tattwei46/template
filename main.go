package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	setupConfig()
	setupLogger()
	r := setupRouter()
	r.Run(getServiceAddr())
}

func getServiceAddr() string {
	if service, ok := serviceInfo[serviceKey]; ok {
		return fmt.Sprintf("%s:%d", service.Host, service.Port)
	}
	return fmt.Sprintf("0.0.0.0:15777")
}

func setupLogger() {
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	v1.GET("/", health)
	return r
}
