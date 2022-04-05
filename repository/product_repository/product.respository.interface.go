package product_repository

import (
	"backend-onboarding/model/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(uuid.UUID) (*entity.Product, []entity.User, error)
	CreateNewProduct(entity.Product) (*entity.Product, error)
	UpdateProductData(entity.Product, string) (*entity.Product, error)
	UpdateCheckedProduct(entity.Product, string) (*entity.Product, error)
	UpdatePublishedProduct(entity.Product, string) (*entity.Product, error)
	DeleteProductById(string) error
}

type productRepository struct {
	Connection *gorm.DB
}

func GetProductRepository(mysqlConn *gorm.DB) ProductRepository {
	return &productRepository{
		Connection: mysqlConn,
	}
}
