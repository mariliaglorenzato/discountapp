package controllers

import (
	"net/http"

	"discountapp/controllers/responses"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type DiscountsController struct {
	getDiscount interfaces.IGetDiscount
}

func NewDiscountsController(
	getDiscount interfaces.IGetDiscount,
) *DiscountsController {
	return &DiscountsController{
		getDiscount: getDiscount,
	}
}

// GetDiscount godoc
// @Summary      Get discount
// @Description  get discount
// @Tags         discounts
// @Accept       json
// @Produce      json
// @Param        product_title   query     string  false  "name search by product title"
// @Param        client_email    query     string  false  "name search by client email"  Format(email)
// @Success      200  {array}   responses.DiscountResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      404  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /discounts [get]
func (c *DiscountsController) Show(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()
	ProductSlugQueryParam := queryParams.Get("product_slug")
	if ProductSlugQueryParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing Query Param: ProductSlug"})
		return
	}
	clientEmailQueryParam := queryParams.Get("client_email")
	if clientEmailQueryParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing Query Param: ClientEmail"})
		return
	}

	input := &inputs.DiscountInput{
		ProductSlug: ProductSlugQueryParam,
		ClientEmail: clientEmailQueryParam,
	}

	ucOutput, err := c.getDiscount.Perform(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses.GetDiscountResponse(ucOutput))
}
