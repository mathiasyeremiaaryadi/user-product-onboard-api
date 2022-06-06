package routes

import (
	"user-product-service/deliveries"

	"github.com/gin-gonic/gin"
)

func InitAppRoute(roleDelivery deliveries.RoleDelivery, userDelivery deliveries.UserDelivery, productDelivery deliveries.ProductDelivery) *gin.Engine {
	router := gin.Default()

	router.NoRoute(deliveries.NoRouteDelivery)

	router.GET("/roles", roleDelivery.GetRolesDelivery)

	router.GET("/users", userDelivery.GetUsersDelivery)
	router.GET("/users/:id", userDelivery.GetUserDelivery)
	router.POST("/users", userDelivery.CreateUserDelivery)
	router.PUT("/users/:id", userDelivery.UpdateUserDelivery)
	router.DELETE("/users/:id", userDelivery.DeleteUserDelivery)

	router.GET("/products", productDelivery.GetProductsDelivery)
	router.GET("/products/:id", productDelivery.GetProductDelivery)
	router.POST("/products", productDelivery.CreateProductDelivery)
	router.PUT("/products/:id", productDelivery.UpdateProductDelivery)
	router.PUT("/products/:id/checked", productDelivery.CheckProductDelivery)
	router.PUT("/products/:id/published", productDelivery.PublishProductDelivery)
	router.DELETE("/products/:id", productDelivery.DeleteProductDelivery)

	return router
}
