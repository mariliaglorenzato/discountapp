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

	getProduct := usecases.NewGetProduct(repository)
	getProducts := usecases.NewGetProducts(repository)
	createProduct := usecases.NewCreateProduct(repository)

	getClients := usecases.NewGetClients(repository)
	createClient := usecases.NewCreateClient(repository)

	getDiscount := usecases.NewGetDiscount(repository)

	// load controllers
	productsController := controllers.NewProductsController(
		getProducts,
		getProduct,
		createProduct,
	)

	clientsController := controllers.NewClientsController(
		getClients,
		createClient,
	)

	discountsController := controllers.NewDiscountsController(
		getDiscount,
	)
	// load routes
	router := gin.Default()
	router.GET("/products", productsController.GetAll)
	router.GET("/product/:title", productsController.Show)
	router.POST("/product", productsController.Create)

	router.GET("/clients", clientsController.GetAll)
	router.POST("/client", clientsController.Create)

	router.GET("/discount/:client_email/*product_title", discountsController.Show)

	router.Run(fmt.Sprintf(":%v", "8080"))
}
