package product_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/model/entity"
	"backend-onboarding/utils"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (product *productUseCase) GetAllProducts() dto.Result {
	productlist, err := product.productRepo.GetAllProducts()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}

	response := []dto.ProductGelAll{}
	for _, product := range productlist {
		resProduct := dto.ProductGelAll{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Status:      product.Status,
		}
		response = append(response, resProduct)
	}

	return utils.ResponseSuccess("ok", nil, response, 200)
}

func (product *productUseCase) GetProductById(id string) dto.Result {
	productUUID, err := uuid.Parse(id)

	if err != nil {
		return utils.ResponseError("Data not found", err, 404)
	}
	productData, userData, err := product.productRepo.GetProductById(productUUID)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}

	maker := dto.Action{}
	checker := dto.Action{}
	signer := dto.Action{}
	for _, user := range userData {
		if user.ID == productData.MakerID {
			maker.ID = user.ID
			maker.Name = user.Name
		}
		if user.ID == productData.CheckerID {
			checker.ID = user.ID
			checker.Name = user.Name
		}
		if user.ID == productData.SignerID {
			signer.ID = user.ID
			signer.Name = user.Name
		}
	}
	response := dto.ProductGetById{
		ID:          productData.ID,
		Name:        productData.Name,
		Description: productData.Description,
		Status:      productData.Status,
		Maker:       maker,
		Checker:     checker,
		Signer:      signer,
	}

	return utils.ResponseSuccess("ok", nil, response, 200)
}

func (product *productUseCase) CreateNewProduct(newProduct dto.Product) dto.Result {
	userInsert := entity.Product{
		ID:          newProduct.ID,
		Name:        newProduct.Name,
		Description: newProduct.Description,
		Status:      "inactive",
		MakerID:     newProduct.MakerID,
	}

	userData, err := product.productRepo.CreateNewProduct(userInsert)

	if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}

	return utils.ResponseSuccess("ok", nil, map[string]interface{}{
		"id": userData.ID}, 201)
}

func (product *productUseCase) UpdateProductData(productUpdate dto.Product, id string) dto.Result {
	productInsert := entity.Product{
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
	}
	_, err := product.productRepo.UpdateProductData(productInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}
	return utils.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUseCase) UpdateCheckProduct(productUpdate dto.Product, id string) dto.Result {
	productInsert := entity.Product{
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
		Status:      "approved",
		CheckerID:   productUpdate.CheckerID,
	}
	_, err := product.productRepo.UpdateCheckedProduct(productInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}
	return utils.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}
func (product *productUseCase) UpdatePublishProduct(productUpdate dto.Product, id string) dto.Result {
	productInsert := entity.Product{
		Name:        productUpdate.Name,
		Description: productUpdate.Description,
		Status:      "active",
		SignerID:    productUpdate.SignerID,
	}
	_, err := product.productRepo.UpdatePublishedProduct(productInsert, id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}
	return utils.ResponseSuccess("ok", nil, map[string]interface{}{"id": id}, 200)
}

func (product *productUseCase) DeleteProductById(id string) dto.Result {

	err := product.productRepo.DeleteProductById(id)

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.ResponseError("Data not found", err, 404)
	} else if err != nil {
		return utils.ResponseError("Internal server error", err, 500)
	}
	return utils.ResponseSuccess("ok", nil, nil, 200)
}
