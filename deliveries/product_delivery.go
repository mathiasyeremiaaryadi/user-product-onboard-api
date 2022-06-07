package deliveries

import (
	"net/http"
	"user-product-service/dto"
	"user-product-service/usecases"

	"github.com/gin-gonic/gin"
)

type ProductDelivery struct {
	productUseCase usecases.ProductUseCase
}

func NewProductDelivery(productUseCase usecases.ProductUseCase) ProductDelivery {
	return ProductDelivery{
		productUseCase: productUseCase,
	}
}

func (productDelivery ProductDelivery) GetProductsDelivery(c *gin.Context) {
	productResponseBody := dto.ResponseBody{}
	products, err := productDelivery.productUseCase.GetProductsUseCase()
	if err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to retrieve products, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, productResponseBody.GetResponseBody())

		return
	}

	productResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(products)

	c.JSON(http.StatusOK, productResponseBody.GetResponseBody())
}

func (productDelivery ProductDelivery) GetProductDelivery(c *gin.Context) {
	productResponseBody := dto.ResponseBody{}
	productID := c.Param("id")

	product, err := productDelivery.productUseCase.GetProductUseCase(productID)
	if err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(404).
			SetError("Failed to retrieve product, product not found").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusNotFound, productResponseBody.GetResponseBody())

		return
	}

	productResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(product)

	c.JSON(http.StatusOK, productResponseBody.GetResponseBody())
}

func (productDelivery ProductDelivery) CreateProductDelivery(c *gin.Context) {
	productPostRequestBody := dto.ProductRequestBody{}
	productResponseBody := dto.ResponseBody{}

	if err := c.ShouldBindJSON(&productPostRequestBody); err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(400).
			SetError("Failed to bind data, data is invalid").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusBadRequest, productResponseBody.GetResponseBody())

		return
	}

	createdProduct, err := productDelivery.productUseCase.CreateProductUseCase(productPostRequestBody)
	if err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to create product, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, productResponseBody.GetResponseBody())

		return
	}

	productResponseBody.
		SetStatus("ok").
		SetStatusCode(201).
		SetError(nil).
		SetData(map[string]interface{}{
			"id": createdProduct.ID,
		})

	c.JSON(http.StatusCreated, productResponseBody.GetResponseBody())
}

func (productDelivery ProductDelivery) UpdateProductDelivery(c *gin.Context) {
	productPutRequestBody := dto.ProductRequestBody{}
	productResponseBody := dto.ResponseBody{}
	productID := c.Param("id")

	if err := c.ShouldBindJSON(&productPutRequestBody); err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(400).
			SetError("Failed to bind data, data is invalid").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusBadRequest, productResponseBody.GetResponseBody())

		return
	}

	updatedProductID, err := productDelivery.productUseCase.UpdateProductUseCase(productPutRequestBody, productID)
	if err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to update product, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, productResponseBody.GetResponseBody())
	}

	productResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(map[string]interface{}{
			"id": updatedProductID,
		})

	c.JSON(http.StatusOK, productResponseBody.GetResponseBody())
}

func (productDelivery ProductDelivery) CheckProductDelivery(c *gin.Context) {
	productPutRequestBody := dto.ProductRequestBody{}
	productResponseBody := dto.ResponseBody{}
	productID := c.Param("id")

	if err := c.ShouldBindJSON(&productPutRequestBody); err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(400).
			SetError("Failed to bind data, data is invalid").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusBadRequest, productResponseBody.GetResponseBody())

		return
	}

	checkedProductID, err := productDelivery.productUseCase.CheckProductUseCase(productPutRequestBody, productID)
	if err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to check product, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, productResponseBody.GetResponseBody())
	}

	productResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(map[string]interface{}{
			"id": checkedProductID,
		})

	c.JSON(http.StatusOK, productResponseBody.GetResponseBody())
}

func (productDelivery ProductDelivery) PublishProductDelivery(c *gin.Context) {
	productPutRequestBody := dto.ProductRequestBody{}
	productResponseBody := dto.ResponseBody{}
	productID := c.Param("id")

	if err := c.ShouldBindJSON(&productPutRequestBody); err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(400).
			SetError("Failed to bind data, data is invalid").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusBadRequest, productResponseBody.GetResponseBody())

		return
	}

	publishedProductID, err := productDelivery.productUseCase.PublishProductUseCase(productPutRequestBody, productID)
	if err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(500).
			SetError("Failed to publish product, server disconnected or error").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusInternalServerError, productResponseBody.GetResponseBody())

		return
	}

	productResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(map[string]interface{}{
			"id": publishedProductID,
		})

	c.JSON(http.StatusOK, productResponseBody.GetResponseBody())
}

func (productDelivery ProductDelivery) DeleteProductDelivery(c *gin.Context) {
	productResponseBody := dto.ResponseBody{}
	productID := c.Param("id")

	err := productDelivery.productUseCase.DeleteProductUseCase(productID)
	if err != nil {
		productResponseBody.
			SetStatus("failed").
			SetStatusCode(404).
			SetError("Failed to delete product, product not found").
			SetData(nil)

		c.AbortWithStatusJSON(http.StatusNotFound, productResponseBody.GetResponseBody())

		return
	}

	productResponseBody.
		SetStatus("ok").
		SetStatusCode(200).
		SetError(nil).
		SetData(nil)

	c.JSON(http.StatusOK, productResponseBody.GetResponseBody())
}
