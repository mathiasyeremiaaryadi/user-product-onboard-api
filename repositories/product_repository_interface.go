package repositories

import "user-product-service/entities"

type ProductRepository interface {
	GetProductsRepository(productList []entities.ProductList) ([]entities.ProductList, error)
	GetProductRepository(productDetail entities.ProductDetail, id int) (entities.ProductDetail, error)
	CreateProductRepository(product entities.Product) (entities.Product, error)
	UpdateProductRepository(product entities.Product, id int) (int, error)
	CheckProductRepository(product entities.Product, id int) (int, error)
	PublishProductRepository(product entities.Product, id int) (int, error)
	DeleteProductRepository(product entities.Product, id int) error
}
