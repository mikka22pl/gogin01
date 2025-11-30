package main

import (
	"gogin01/middleware"
	"io"

	"github.com/gin-gonic/gin"
)

func main02() {
	router := gin.Default()

	// middleware

	// router.Use(middleware.Authenticate)

	/*auth := gin.BasicAuth(gin.Accounts{
		"user":  "pass",
		"user2": "pass2",
		"user3": "pass3",
	})*/

	router.GET("/getData", getData)
	router.GET("/getData1", middleware.AddHeader, getData1)
	router.GET("/getData2", middleware.Authenticate, middleware.AddHeader, getData2)
	// router.GET("/getQueryString", getQueryString)
	// router.GET("/getParams/name/:name/age/:age", getParams)
	router.POST("/data", dataPost)

	/*
		admin := router.Group("/admin", auth)
		admin := router.Group("/admin", middleware.Authenticate)
		{
			// localhost:9091/admin/getQueryString
			admin.GET("/getQueryString", getQueryString)
		}
	*/

	client := router.Group("/client")
	{
		// localhost:9091/client/getParams/name/...
		client.GET("/getParams/name/:name/age/:age", getParams)
	}
	router.Run() // listen and serve on 0.0.0.0:8080

	// http.ListenAndServe(":9090", router)
	/*server := &http.Server{
		Addr:         ":9091",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()*/
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am GetData Method",
	})
}
func getData1(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am GetData 1 Method",
	})
}
func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "hi I am GetData 2 Method",
	})
}

// getQueryString?name=Mikka&age=30
func getQueryString(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "hi I am getQueryString method in Gin framework",
		"name": name,
		"age":  age,
	})
}

// getParams/name/Mikka/age/30
func getParams(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data": "hi I am getParams method in Gin framework",
		"name": name,
		"age":  age,
	})
}

func dataPost(c *gin.Context) {
	body := c.Request.Body
	value, _ := io.ReadAll(body)

	c.JSON(201, gin.H{
		"data": "hi I am POST method from GIN framework",
		"body": string(value),
	})
}
