package controllers

import (
	"net/http"

	"discountapp/controllers/payloads"
	"discountapp/controllers/responses"
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
// @Success      200  {array}   responses.ProductResponse
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

	ctx.IndentedJSON(http.StatusOK, responses.GetProductsResponse(ucOutput))
}

// AddProduct godoc
// @Summary      Add a product
// @Description  add by json product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      payloads.ProductPayload  true  "Add product"
// @Success      200      {object}  responses.ProductResponse
// @Failure      400      {object}  responses.ErrorResponse
// @Failure      404      {object}  responses.ErrorResponse
// @Failure      500      {object}  responses.ErrorResponse
// @Router       /products [post]
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
// @Param        slug   path      string  true  "Product Slug"
// @Success      200  {object}  responses.ProductResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      404  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
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
