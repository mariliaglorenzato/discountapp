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

// GetAll godoc
// @Summary      List products
// @Description  get products
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {array}   payloads.ProductPayload
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      404  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /products [get]
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

// Show godoc
// @Summary      Show an account
// @Description  get product by slug
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        slug   path      string  true  "Title Slug"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /products/{slug} [get]
func (c *ProductsController) Show(ctx *gin.Context) {
	titleSlugParam := ctx.Param("slug")
	if titleSlugParam == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Missing Param: Slug"})
		return
	}

	input := inputs.ProductBySlugInput{Slug: titleSlugParam}

	ucOutput, err := c.getProduct.Perform(&input)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, ucOutput)
}
