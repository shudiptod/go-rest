package routes

import (
	"pcshop/app/handlers"

	"github.com/gin-gonic/gin"
)



func SetupRoutes(router *gin.Engine){
	v1 := router.Group("/v1")
	{
		v1.GET("/products",handlers.GetProducts)
		v1.POST("products",handlers.CreateProduct)
	}
}