package controllers

import (
	"net/http"

	"discountapp/controllers/payloads"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	getProducts   interfaces.IGetProducts
	getProduct    interfaces.IGetProduct
	createProduct interfaces.ICreateProduct
}

func NewProductsController(
	getProducts interfaces.IGetProducts,
	getProduct interfaces.IGetProduct,
	createProduct interfaces.ICreateProduct,
) *ProductsController {
	return &ProductsController{
		getProducts:   getProducts,
		getProduct:    getProduct,
		createProduct: createProduct,
	}
}

func (c *ProductsController) GetAll(ctx *gin.Context) {
	ucOutput, err := c.getProducts.Perform()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, ucOutput)
}

func (c *ProductsController) Create(ctx *gin.Context) {
	var payload *payloads.ProductPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productInput := &inputs.ProductInput{
		Title:       payload.Title,
		Description: payload.Description,
		Price:       payload.Price,
	}

	ucOutput, err := c.createProduct.Perform(productInput)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": ucOutput})
}

func (c *ProductsController) Show(ctx *gin.Context) {
	titleParam := ctx.Param("title")
	if titleParam == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing Param: Title"})
		return
	}

	input := inputs.ProductByTitleInput{Title: titleParam}

	ucOutput, err := c.getProduct.Perform(&input)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, ucOutput)
}
