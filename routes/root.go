package routes

import "github.com/gin-gonic/gin"

func GetRootHandler() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Tricking API!",
		})
	}

	return gin.HandlerFunc(fn)
}
