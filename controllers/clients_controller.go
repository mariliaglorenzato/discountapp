package controllers

import (
	"net/http"
	"time"

	"discountapp/controllers/payloads"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"

	"github.com/gin-gonic/gin"
)

type ClientsController struct {
	getClient    interfaces.IGetClients
	createClient interfaces.ICreateClient
}

func NewClientsController(
	getClient interfaces.IGetClients,
	createClient interfaces.ICreateClient,
) *ClientsController {
	return &ClientsController{
		getClient:    getClient,
		createClient: createClient,
	}
}

func (c *ClientsController) GetAll(ctx *gin.Context) {
	ucOutput, err := c.getClient.Perform()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, ucOutput)
}

func (c *ClientsController) Create(ctx *gin.Context) {
	var payload *payloads.ClientPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parsedBirthDate, err := time.Parse("02/01/2006", payload.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientInput := &inputs.ClientInput{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		BirthDate: parsedBirthDate,
	}

	ucOutput, err := c.createClient.Perform(clientInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": ucOutput})
}
