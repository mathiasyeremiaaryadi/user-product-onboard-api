package usecases

import (
	"user-product-service/dto"
	"user-product-service/entities"
)

type ProductUseCase interface {
	GetProductsUseCase() ([]entities.ProductList, error)
	GetProductUseCase(id string) (entities.ProductDetail, error)
	CreateProductUseCase(productPostRequestBody dto.ProductRequestBody) (entities.Product, error)
	UpdateProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error)
	CheckProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error)
	PublishProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error)
	DeleteProductUseCase(id string) error
}
