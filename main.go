package main

import (
	"gogin01/logger"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
)

func main() {
	router := gin.New()

	// 2025/11/30 12:25:26 endpoint formatted information is /getData GET main.getData3 3
	gin.DebugPrintRouteFunc = func(httpMethod string, absolutePath string, handlerName string, nuHandlers int) {
		log.Printf("endpoint formatted information is %v %v %v %v\n", absolutePath, httpMethod, handlerName, nuHandlers)
	}

	// colorable the logs in console
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	f, _ := os.Create("ginLogging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// router.Use(gin.LoggerWithFormatter(logger.FormatLogs))
	router.Use(gin.LoggerWithFormatter(logger.FormatLogsJson))

	router.GET("/getData", getData3)
	router.Run()
}
func getData3(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am GetData Method",
	})
}
