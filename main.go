package main

import (
	"discountapp/config"
	"discountapp/controllers"
	"discountapp/docs"
	"discountapp/persistence"
	"discountapp/usecases"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	config.SetConfigs(viper.GetString(config.ServerEnv))
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
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{

		products := v1.Group("/products")
		{
			products.GET(":slug", productsController.Show)
			products.GET("", productsController.GetAll)
			products.POST("", productsController.Create)
		}
		clients := v1.Group("/clients")
		{
			clients.GET("", clientsController.GetAll)
			clients.POST("", clientsController.Create)
		}
		discounts := v1.Group("/discounts")
		{
			discounts.GET("", discountsController.Show)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
