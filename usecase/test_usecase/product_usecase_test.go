package test_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/model/entity"
	"backend-onboarding/repository/test_repository"
	"backend-onboarding/usecase/product_usecase"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var productRepository = test_repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productUc = product_usecase.ProductUsecaseTest{ProductRepo: &productRepository}

func TestGetAllProductsSuccess(t *testing.T) {
	expected := []entity.Product{}
	productRepository.Mock.On("GetAllProducts").Return(expected, nil)
	result, err := productUc.ProductRepo.GetAllProducts()

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestGetProductByIdNotFound(t *testing.T) {
	productRepository.Mock.On("GetProductById", "1").Return(nil, errors.New("Data not found"))

	result, err := productUc.ProductRepo.GetProductById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestGetProductByIdSuccess(t *testing.T) {
	expected := &entity.Product{}

	productRepository.Mock.On("GetProductById", "2").Return(expected, nil)

	result, err := productUc.ProductRepo.GetProductById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestUpdateProductNotFound(t *testing.T) {
	productRepository.Mock.On("GetProductById", "1").Return(nil, errors.New("Data not found"))

	result, err := productUc.ProductRepo.GetProductById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestUpdateProductSuccess(t *testing.T) {
	expected := &entity.Product{}

	productRepository.Mock.On("UpdateProductData", "2").Return(expected, nil)

	result, err := productUc.ProductRepo.UpdateProductData("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewProductSuccess(t *testing.T) {
	request := dto.Product{
		Name:        "New Product",
		Description: "Test New Product",
	}

	expected := &entity.Product{}

	productRepository.Mock.On("CreateNewProduct").Return(expected, nil)

	result, err := productUc.ProductRepo.CreateNewProduct(request)

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}

func TestCreateNewProductSFailed(t *testing.T) {
	request := dto.Product{
		Name:        "New Product",
		Description: "Test New Product",
	}

	expected := &entity.Product{}

	productRepository.Mock.On("CreateNewProduct").Return(nil, errors.New("Failed create product"))

	result, err := productUc.ProductRepo.CreateNewProduct(request)

	assert.NotEqual(t, expected, result)
	assert.NotNil(t, err)
	assert.Nil(t, result)

}

func TestDeleteProductFailed(t *testing.T) {
	productRepository.Mock.On("DeleteProductById", "1").Return(nil, errors.New("Data not found"))

	result, err := productUc.ProductRepo.DeleteProductById("1")
	assert.Nil(t, result)
	assert.NotNil(t, err)

}

func TestDeleteProductSuccess(t *testing.T) {
	expected := &entity.Product{}

	productRepository.Mock.On("DeleteProductById", "2").Return(expected, nil)

	result, err := productUc.ProductRepo.DeleteProductById("2")

	assert.Equal(t, expected, result)
	assert.NotNil(t, result)
	assert.Nil(t, err)

}
