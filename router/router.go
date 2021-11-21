package router

import (
	"github.com/gin-gonic/gin"
	"project-kp/router/endpoint"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

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
