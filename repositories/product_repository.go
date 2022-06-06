package repositories

import (
	"errors"
	"user-product-service/entities/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (productRepository ProductRepository) GetProductsRepository(productList []models.ProductList) ([]models.ProductList, error) {
	if err := productRepository.db.Model(&models.Product{}).Find(&productList).Error; err != nil {
		return productList, err
	}

	return productList, nil
}

func (productRepository ProductRepository) GetProductRepository(productDetail models.ProductDetail, id int) (models.ProductDetail, error) {
	if err := productRepository.db.Preload("Maker").Preload("Checker").Preload("Signer").Model(&models.Product{}).First(&productDetail, id).Error; err != nil {
		return productDetail, err
	}

	return productDetail, nil
}

func (productRepository ProductRepository) CreateProductRepository(product models.Product) (models.Product, error) {
	if err := productRepository.db.Omit("ID", "Status", "CheckerID", "SignerID").Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (productRepository ProductRepository) UpdateProductRepository(product models.Product, id int) (int, error) {
	if result := productRepository.db.Where("id = ?", id).Updates(&product); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (productRepository ProductRepository) CheckProductRepository(product models.Product, id int) (int, error) {
	if result := productRepository.db.Where("id = ?", id).Updates(&product); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (productRepository ProductRepository) PublishProductRepository(product models.Product, id int) (int, error) {
	if result := productRepository.db.Where("id = ?", id).Updates(&product); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (productRepository ProductRepository) DeleteProductRepository(product models.Product, id int) error {
	if err := productRepository.db.First(&product, id).Error; err != nil {
		return err
	}

	if err := productRepository.db.Delete(&product, id).Error; err != nil {
		return err
	}

	return nil
}
