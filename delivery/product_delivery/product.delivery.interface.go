package product_delivery

import (
	"backend-onboarding/usecase/product_usecase"
	"github.com/gin-gonic/gin"
)

type ProductDelivery interface {
	GetAllProducts(*gin.Context)
	CreateNewProduct(*gin.Context)
	UpdateProductById(*gin.Context)
	DeleteProductById(c *gin.Context)
	GetProductById(c *gin.Context)
	UpdateCheckedProduct(c *gin.Context)
	UpdatePublishedProduct(c *gin.Context)
}

type productDelivery struct {
	productUseCase product_usecase.ProductUseCase
}

func GetProductDelivery(productUseCase product_usecase.ProductUseCase) ProductDelivery {
	return &productDelivery{
		productUseCase: productUseCase,
	}
}
