package porducts

import (
	"net/http"
	Service "product-api/app/services/ServiceProduct"
	"product-api/app/models/products"
	"product-api/app/utils/errors"

	"github.com/gin-gonic/gin"
)

func RegisterProduct(ctx *gin.Context) {
	var pro products.Product
	if err := ctx.ShouldBindJSON(&pro); err != nil {
		err := errors.BadRequestError(" invalid Json")
		ctx.JSON(err.Status, err)
		return
	}
	newPro, newErr := Service.AddProduct(pro)
	if newErr != nil {
		err := errors.InternalServerError(" cannot register product")
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusCreated,
		"data":   newPro,
	})

}
func GetProducts(ctx *gin.Context) {
	pro := Service.GetProduct()
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusAccepted,
		"data":   pro,
	})
}
