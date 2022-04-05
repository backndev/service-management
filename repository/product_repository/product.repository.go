package product_repository

import (
	"backend-onboarding/model/entity"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (productRepo *productRepository) GetAllProducts() ([]entity.Product, error) {
	products := []entity.Product{}
	err := productRepo.Connection.Model(&entity.Product{}).Scan(&products).Error
	if err != nil {
		return nil, err
	}

	if len(products) <= 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return products, nil
}

func (productRepo *productRepository) GetProductById(id uuid.UUID) (*entity.Product, []entity.User, error) {
	product := entity.Product{}
	users := []entity.User{}

	if err := productRepo.Connection.Where("id = ?", id).Find(&product).Error; err != nil {
		return nil, nil, err
	}

	if (entity.Product{}) == product {
		return nil, nil, gorm.ErrRecordNotFound
	}

	err := productRepo.Connection.Where("id IN ?", []uuid.UUID{product.MakerID, product.CheckerID, product.SignerID}).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	return &product, users, nil
}

func (productRepo *productRepository) CreateNewProduct(product entity.Product) (*entity.Product, error) {
	product.ID = uuid.New()

	if err := productRepo.Connection.Create(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (productRepo *productRepository) UpdateProductData(product entity.Product, id string) (*entity.Product, error) {
	result := productRepo.Connection.Model(&product).Where("id = ?", id).Updates(&product)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &product, nil
}

func (productRepo *productRepository) UpdateCheckedProduct(product entity.Product, id string) (*entity.Product, error) {
	result := productRepo.Connection.Model(&product).Where("id = ?", id).Updates(&product)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &product, nil
}

func (productRepo *productRepository) UpdatePublishedProduct(product entity.Product, id string) (*entity.Product, error) {
	result := productRepo.Connection.Model(&product).Where("id = ?", id).Updates(&product)
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &product, nil
}

func (productRepo *productRepository) DeleteProductById(id string) error {
	sql := "DELETE FROM products"
	sql = fmt.Sprintf("%s WHERE id = '%s'", sql, id)
	if err := productRepo.Connection.Raw(sql).Scan(entity.Product{}).Error; err != nil {
		return err
	}

	return nil
}
