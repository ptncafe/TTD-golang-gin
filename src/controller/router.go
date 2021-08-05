package controller

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", HomeController)
	router.GET("/ping", HomeController)
	api := router.Group("/api")
	{
		v1 := api.Group("/v1/shop")
		{
			shop := v1.Group("/v1/shop")
			{
				shop.POST("/update-name", UpdateNameController)
			}
		}
	}
}
