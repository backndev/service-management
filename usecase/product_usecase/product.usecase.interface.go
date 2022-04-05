package product_usecase

import (
	"backend-onboarding/model/dto"
	"backend-onboarding/repository/product_repository"
)

type ProductUseCase interface {
	GetAllProducts() dto.Result
	GetProductById(string) dto.Result
	CreateNewProduct(dto.Product) dto.Result
	UpdateProductData(dto.Product, string) dto.Result
	UpdatePublishProduct(dto.Product, string) dto.Result
	UpdateCheckProduct(dto.Product, string) dto.Result
	DeleteProductById(string) dto.Result
}

type productUseCase struct {
	productRepo product_repository.ProductRepository
}

func GetProductUseCase(productRepository product_repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		productRepo: productRepository,
	}
}
