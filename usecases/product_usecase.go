package usecases

import (
	"strconv"
	"user-product-service/entities/dto"
	"user-product-service/entities/models"
	"user-product-service/repositories"
)

type ProductUseCase struct {
	productRepository repositories.ProductRepository
}

func NewProductUseCase(productRepository repositories.ProductRepository) ProductUseCase {
	return ProductUseCase{
		productRepository: productRepository,
	}
}

func (productUseCase ProductUseCase) GetProductsUseCase() ([]models.ProductList, error) {
	var productList []models.ProductList

	products, err := productUseCase.productRepository.GetProductsRepository(productList)
	if err != nil {
		return products, err
	}

	return products, err
}

func (productUseCase ProductUseCase) GetProductUseCase(id string) (models.ProductDetail, error) {
	var productDetail models.ProductDetail

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productDetail, err
	}

	product, err := productUseCase.productRepository.GetProductRepository(productDetail, productID)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (productUseCase ProductUseCase) CreateProductUseCase(productPostRequestBody dto.ProductRequestBody) (models.Product, error) {
	productData := models.Product{
		Name:        productPostRequestBody.Name,
		Description: productPostRequestBody.Description,
		MakerID:     2,
	}

	createdProduct, err := productUseCase.productRepository.CreateProductRepository(productData)
	if err != nil {
		return createdProduct, err
	}

	return createdProduct, nil
}

func (productUseCase ProductUseCase) UpdateProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error) {
	productData := models.Product{
		Name:        productPutRequestBody.Name,
		Description: productPutRequestBody.Description,
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productID, err
	}

	updatedProductID, err := productUseCase.productRepository.UpdateProductRepository(productData, productID)
	if err != nil {
		return updatedProductID, err
	}

	return updatedProductID, nil
}

func (productUseCase ProductUseCase) CheckProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error) {
	productData := models.Product{
		Name:        productPutRequestBody.Name,
		Description: productPutRequestBody.Description,
		Status:      "approved",
		CheckerID:   4,
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productID, err
	}

	checkedProductID, err := productUseCase.productRepository.CheckProductRepository(productData, productID)
	if err != nil {
		return checkedProductID, err
	}

	return checkedProductID, nil
}

func (productUseCase ProductUseCase) PublishProductUseCase(productPutRequestBody dto.ProductRequestBody, id string) (int, error) {
	productData := models.Product{
		Name:        productPutRequestBody.Name,
		Description: productPutRequestBody.Description,
		Status:      "active",
		SignerID:    6,
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		return productID, err
	}

	publishedProductID, err := productUseCase.productRepository.PublishProductRepository(productData, productID)
	if err != nil {
		return publishedProductID, err
	}

	return publishedProductID, nil
}

func (productUseCase ProductUseCase) DeleteProductUseCase(id string) error {
	var deletedProduct models.Product

	productID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = productUseCase.productRepository.DeleteProductRepository(deletedProduct, productID)
	if err != nil {
		return err
	}

	return nil
}
