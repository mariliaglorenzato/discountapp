package controllers

import (
	"net/http"

	"discountapp/controllers/params"
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

func (c *DiscountsController) Show(ctx *gin.Context) {
	var params params.DiscountParams

	if err := ctx.ShouldBind(&params); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input := inputs.DiscountInput{
		ProductTitle: params.ProductTitle,
		ClientEmail:  params.ClientEmail,
	}

	ucOutput, err := c.getDiscount.Perform(&input)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, ucOutput)
}
