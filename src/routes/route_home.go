package routes

import "github.com/gin-gonic/gin"

// MainRoute ...
func MainRoute(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
