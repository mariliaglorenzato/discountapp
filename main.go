package main

import (
	"fmt"

	"discountapp/controllers"
	"discountapp/persistence"
	"discountapp/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	// load database
	db := persistence.NewPersistence()
	// load repository
	repository := persistence.NewRepository(db.DB)

	// load use cases
	getProducts := usecases.NewGetProducts(repository)
	createProduct := usecases.NewCreateProduct(repository)

	getClients := usecases.NewGetClients(repository)
	createClient := usecases.NewCreateClient(repository)

	// load controllers
	productsController := controllers.NewProductsController(
		getProducts,
		createProduct,
	)

	clientsController := controllers.NewClientsController(
		getClients,
		createClient,
	)

	// load routes
	router := gin.Default()
	router.GET("/products", productsController.GetAll)
	router.POST("/product", productsController.Create)

	router.GET("/clients", clientsController.GetAll)
	router.POST("/client", clientsController.Create)

	router.Run(fmt.Sprintf(":%v", "8080"))
}
