package deliveries

import (
	"net/http"
	"user-product-service/entities/dto"

	"github.com/gin-gonic/gin"
)

func NoRouteDelivery(c *gin.Context) {
	notFoundResponseBody := dto.ResponseBody{}

	notFoundResponseBody.
		SetStatus("failed").
		SetStatusCode(404).
		SetError("Resource not found").
		SetData(nil)

	c.JSON(http.StatusNotFound, notFoundResponseBody.GetResponseBody())
}
