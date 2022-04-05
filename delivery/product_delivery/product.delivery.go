package product_delivery

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (product *productDelivery) GetAllProducts(c *gin.Context) {
	result := product.productUseCase.GetAllProducts()
	if result.Status != "ok" {
		c.JSON(result.StatusCode, result)
		return
	}
	c.JSON(result.StatusCode, result)
}

func (product *productDelivery) GetProductById(c *gin.Context) {
	id := c.Param("id")
	result := product.productUseCase.GetProductById(id)

	if result.StatusCode == http.StatusNotFound {
		c.JSON(http.StatusOK, result)
		return
	}
	if result.Status != "ok" {
		c.JSON(result.StatusCode, result)
		return
	}
	c.JSON(http.StatusOK, result)
}

func (product *productDelivery) CreateNewProduct(c *gin.Context) {
	userId, _ := c.Get("user_id")
	userMakerId, _ := userId.(string)
	userMakerUUID, err := uuid.Parse(userMakerId)

	if err != nil {
		errorRes := utils.ResponseError("Data not found", err, 404)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	request := dto.Product{
		MakerID: userMakerUUID,
	}

	if err := c.ShouldBindJSON(&request); err != nil {

		errorRes := utils.ResponseError("Invalid Input", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return

	}
	result := product.productUseCase.CreateNewProduct(request)

	if result.Status != "ok" {
		c.JSON(result.StatusCode, result)
		return
	}

	c.JSON(result.StatusCode, result)
}

func (product *productDelivery) UpdateProductById(c *gin.Context) {
	id := c.Param("id")
	request := dto.Product{}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := utils.ResponseError("Invalid Input", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	result := product.productUseCase.UpdateProductData(request, id)
	if result.Status != "ok" {
		c.JSON(result.StatusCode, result)
		return
	}

	c.JSON(result.StatusCode, result)
}

func (product *productDelivery) UpdateCheckedProduct(c *gin.Context) {
	id := c.Param("id")
	userId, _ := c.Get("user_id")
	userCheckedId, _ := userId.(string)
	userCheckedUUID, err := uuid.Parse(userCheckedId)

	if err != nil {
		errorRes := utils.ResponseError("Data not found", err, 404)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	request := dto.Product{
		CheckerID: userCheckedUUID,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := utils.ResponseError("Invalid Input", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := product.productUseCase.UpdateCheckProduct(request, id)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}
func (product *productDelivery) UpdatePublishedProduct(c *gin.Context) {
	id := c.Param("id")
	userId, _ := c.Get("user_id")
	userSignedId, _ := userId.(string)
	userSignedUUID, err := uuid.Parse(userSignedId)

	if err != nil {
		errorRes := utils.ResponseError("Data not found", err, 404)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}
	request := dto.Product{
		SignerID: userSignedUUID,
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errorRes := utils.ResponseError("Invalid Input", err, 400)
		c.JSON(errorRes.StatusCode, errorRes)
		return
	}

	response := product.productUseCase.UpdatePublishProduct(request, id)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (product *productDelivery) DeleteProductById(c *gin.Context) {
	id := c.Param("id")
	response := product.productUseCase.DeleteProductById(id)
	if response.Status != "ok" {
		c.JSON(response.StatusCode, response)
		return
	}
	c.JSON(response.StatusCode, response)
}
