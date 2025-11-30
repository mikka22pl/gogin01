package middleware

import "github.com/gin-gonic/gin"

func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Token") == "auth") {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Token not present",
		})
		return
	}
	c.Next()
}

func AddHeader(c *gin.Context) {
	c.Writer.Header().Set("Key", "Value")
	c.Next()
}

/*
func Authenticate() gin.HandlerFunc {
	// Write custom login to be applied before my middleware is executed

	return func(c *gin.Context) {
		if !(c.Request.Header.Get("Token") == "auth") {
			c.AbortWithStatusJSON(500, gin.H{
				"Message": "Token not present",
			})
			return
		}
		c.Next()
	}
}
*/
