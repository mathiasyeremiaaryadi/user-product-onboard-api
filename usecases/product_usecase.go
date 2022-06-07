package usecases

import (
	"strconv"
	"user-product-service/entities"
	"user-product-service/entities/dto"
	"user-product-service/repositories"
)

type productUseCase struct {
	productRepository repositories.ProductRepository
}

func NewProductUseCase(productRepository repositories.ProductRepository) ProductUseCase {
	return productUseCase{
		productRepository: productRepository,
	}
}

func (productUC productUseCase) GetProductsUseCase() ([]entities.ProductList, error) {
	var productList []entities.ProductList

	products, err := productUC.productRepository.GetProductsRepository(productList)
	if err != nil {
		return products, err
	}

	return products, err
}

func (productUC productUseCase) GetProductUseCase(id string) (entities.ProductDetail, error) {
	var productDetail entities.ProductDetail

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productDetail, err
	}

	product, err := productUC.productRepository.GetProductRepository(productDetail, productID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (productUC productUseCase) CreateProductUseCase(productPostRequestBody dto.ProductRequestBody) (entities.Product, error) {
	productData := entities.Product{
		Name:        productPostRequestBody.Name,
		Description: productPostRequestBody.Description,
		MakerID:     2,
	}

	createdProduct, err := productUC.productRepository.CreateProductRepository(productData)
	if err != nil {
		return createdProduct, err
	}

	return createdProduct, nil
}

func (productUC productUseCase) UpdateProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error) {
	productData := entities.Product{
		Name:        productPutRequestBody.Name,
		Description: productPutRequestBody.Description,
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productID, err
	}

	updatedProductID, err := productUC.productRepository.UpdateProductRepository(productData, productID)
	if err != nil {
		return updatedProductID, err
	}

	return updatedProductID, nil
}

func (productUC productUseCase) CheckProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error) {
	productData := entities.Product{
		Name:        productPutRequestBody.Name,
		Description: productPutRequestBody.Description,
		Status:      "approved",
		CheckerID:   4,
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productID, err
	}

	checkedProductID, err := productUC.productRepository.CheckProductRepository(productData, productID)
	if err != nil {
		return checkedProductID, err
	}

	return checkedProductID, nil
}

func (productUC productUseCase) PublishProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error) {
	productData := entities.Product{
		Name:        productPutRequestBody.Name,
		Description: productPutRequestBody.Description,
		Status:      "active",
		SignerID:    6,
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productID, err
	}

	publishedProductID, err := productUC.productRepository.PublishProductRepository(productData, productID)
	if err != nil {
		return publishedProductID, err
	}

	return publishedProductID, nil
}

func (productUC productUseCase) DeleteProductUseCase(id string) error {
	var deletedProduct entities.Product

	productID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = productUC.productRepository.DeleteProductRepository(deletedProduct, productID)
	if err != nil {
		return err
	}

	return nil
}
