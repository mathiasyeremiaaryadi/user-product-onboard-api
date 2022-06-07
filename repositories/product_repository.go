package repositories

import (
	"errors"
	"user-product-service/entities"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{
		db: db,
	}
}

func (productRepo productRepository) GetProductsRepository(productList []entities.ProductList) ([]entities.ProductList, error) {
	if err := productRepo.db.Model(&entities.Product{}).Find(&productList).Error; err != nil {
		return productList, err
	}

	return productList, nil
}

func (productRepo productRepository) GetProductRepository(productDetail entities.ProductDetail, id int) (entities.ProductDetail, error) {
	if err := productRepo.db.Preload("Maker").Preload("Checker").Preload("Signer").Model(&entities.Product{}).First(&productDetail, id).Error; err != nil {
		return productDetail, err
	}

	return productDetail, nil
}

func (productRepo productRepository) CreateProductRepository(product entities.Product) (entities.Product, error) {
	if err := productRepo.db.Omit("ID", "Status", "CheckerID", "SignerID").Create(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (productRepo productRepository) UpdateProductRepository(product entities.Product, id int) (int, error) {
	if result := productRepo.db.Where("id = ?", id).Updates(&product); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (productRepo productRepository) CheckProductRepository(product entities.Product, id int) (int, error) {
	if result := productRepo.db.Where("id = ?", id).Updates(&product); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (productRepo productRepository) PublishProductRepository(product entities.Product, id int) (int, error) {
	if result := productRepo.db.Where("id = ?", id).Updates(&product); result.RowsAffected == 0 {
		return id, errors.New("failed")
	}

	return id, nil
}

func (productRepo productRepository) DeleteProductRepository(product entities.Product, id int) error {
	if err := productRepo.db.First(&product, id).Error; err != nil {
		return err
	}

	if err := productRepo.db.Delete(&product, id).Error; err != nil {
		return err
	}

	return nil
}
