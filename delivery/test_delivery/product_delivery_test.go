package test_delivery

import (
	"backend-onboarding/delivery/product_delivery"
	"backend-onboarding/model/dto"
	"backend-onboarding/usecase/mock_usecase"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productUsecase = mock_usecase.ProductUsecaseMock{Mock: mock.Mock{}}
var productDel = product_delivery.ProductDeliveryTest{ProductUsecase: &productUsecase}

func TestGetAllProductsSuccess(t *testing.T) {

	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       []dto.ProductGelAll{},
	}

	productUsecase.Mock.On("GetAllProducts").Return(expected)

	result := productDel.ProductUsecase.GetAllProducts()

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestGetProductByIdSuccess(t *testing.T) {

	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       dto.ProductGetById{},
	}

	productUsecase.Mock.On("GetProductById", "1").Return(expected)

	result := productDel.ProductUsecase.GetProductById("1")

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestGetProductByIdNotFound(t *testing.T) {

	expected := dto.Response{
		StatusCode: 404,
		Status:     "Data not found",
		Error:      errors.New("Record not found"),
		Data:       nil,
	}

	productUsecase.Mock.On("GetProductById", "2").Return(expected)

	result := productDel.ProductUsecase.GetProductById("2")

	assert.Equal(t, expected.Error, result.Error)
	assert.Nil(t, result.Data)
	assert.NotNil(t, result.Error)
}
func TestDeleteProductSuccess(t *testing.T) {
	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       map[string]interface{}{"id": "1"},
	}

	productUsecase.Mock.On("DeleteProductById", "1").Return(expected)

	result := productDel.ProductUsecase.DeleteProductById("1")

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestDeleteProductFailed(t *testing.T) {

	expected := dto.Response{
		StatusCode: 404,
		Status:     "Data not found",
		Error:      errors.New("Record not found"),
		Data:       nil,
	}

	productUsecase.Mock.On("DeleteProductById", "2").Return(expected)

	result := productDel.ProductUsecase.DeleteProductById("2")

	assert.Equal(t, expected.Error, result.Error)
	assert.Nil(t, result.Data)
	assert.NotNil(t, result.Error)
}

func TestUpdateProductDataSuccess(t *testing.T) {
	expected := dto.Response{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       map[string]interface{}{"id": "1"},
	}

	productUsecase.Mock.On("UpdateProductData", "1").Return(expected)

	result := productDel.ProductUsecase.UpdateProductData("1")

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
func TestUpdateProductDataFailed(t *testing.T) {

	expected := dto.Response{
		StatusCode: 404,
		Status:     "Data not found",
		Error:      errors.New("Record not found"),
		Data:       nil,
	}

	productUsecase.Mock.On("UpdateProductData", "2").Return(expected)

	result := productDel.ProductUsecase.UpdateProductData("2")

	assert.Equal(t, expected.Error, result.Error)
	assert.Nil(t, result.Data)
	assert.NotNil(t, result.Error)
}

func TestCreateNewProductSuccess(t *testing.T) {
	expected := dto.Result{
		StatusCode: 200,
		Status:     "ok",
		Error:      nil,
		Data:       map[string]interface{}{"id": "1"},
	}

	productUsecase.Mock.On("CreateNewProduct").Return(expected)

	result := productDel.ProductUsecase.CreateNewProduct(dto.Product{Name: "test", Description: "test"})

	assert.Equal(t, expected, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Data)
}
