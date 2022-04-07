package routes

import (
	"net/http"
	 products "product-api/app/controller/ProductsController"
	"github.com/gin-gonic/gin"
)

func Url_maps() {
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Is Healthy")
	})
	v1 := r.Group("/api/v1/products")
	{
		v1.POST("/register",products.RegisterProduct)
		v1.GET("/", products.GetProducts)
	}
}
