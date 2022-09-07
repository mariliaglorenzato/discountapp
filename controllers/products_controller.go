package controllers

import (
	"net/http"

	"discountapp/controllers/payloads"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	getProduct    interfaces.IGetProducts
	createProduct interfaces.ICreateProduct
}

func NewProductsController(
	getProduct interfaces.IGetProducts,
	createProduct interfaces.ICreateProduct,
) *ProductsController {
	return &ProductsController{
		getProduct:    getProduct,
		createProduct: createProduct,
	}
}

func (c *ProductsController) GetAll(ctx *gin.Context) {
	ucOutput, err := c.getProduct.Perform()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, ucOutput)
}

func (c *ProductsController) Create(ctx *gin.Context) {
	var payload *payloads.ProductPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productInput := &inputs.ProductInput{
		Title:       payload.Title,
		Description: payload.Description,
		Price:       payload.Price,
	}

	ucOutput, err := c.createProduct.Perform(productInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ucOutput})
}
