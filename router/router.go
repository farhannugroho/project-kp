package router

import (
	"net/http"
	"project-kp/router/endpoint"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":   "OK",
			"base_url": "https://warehouse-kp.herokuapp.com/v1/",
		})
	})

	v1 := r.Group("/v1")
	{
		// Uom
		uom := v1.Group("uom")
		{
			uom.GET("", endpoint.GetAllUom)
			uom.POST("", endpoint.CreateUom)
			uom.DELETE("/:id", endpoint.DeleteUom)
		}
	}

	return r
}
