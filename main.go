package main

import (
	"user-product-service/config"
	"user-product-service/database"
	"user-product-service/deliveries"
	"user-product-service/repositories"
	"user-product-service/routes"
	"user-product-service/usecases"
)

func main() {
	config.InitConfig()

	mysqldb := database.MySQLDatabase{}
	mysqldb.SetDBConnection()
	mysqldb.SetDBSeed()

	roleRepository := repositories.NewRoleRepository(mysqldb.GetDBConnection())
	userRepository := repositories.NewUserRepository(mysqldb.GetDBConnection())
	productRepository := repositories.NewProductRepository(mysqldb.GetDBConnection())

	roleUseCase := usecases.NewRoleUseCase(roleRepository)
	userUseCase := usecases.NewUserUseCase(userRepository)
	productUseCase := usecases.NewProductUseCase(productRepository)

	roleDelivery := deliveries.NewRoleDelivery(roleUseCase)
	userDelivery := deliveries.NewUserDelivery(userUseCase)
	productDelivery := deliveries.NewProductDelivery(productUseCase)

	router := routes.InitAppRoute(roleDelivery, userDelivery, productDelivery)
	router.Run(config.ENV["APP_HOST"])
}
