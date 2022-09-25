package controllers

import (
	"net/http"
	"time"

	"discountapp/controllers/payloads"
	"discountapp/controllers/responses"
	"discountapp/usecases/inputs"
	"discountapp/usecases/interfaces"
	"discountapp/utils"

	"github.com/gin-gonic/gin"
)

type ClientsController struct {
	getClients   interfaces.IGetClients
	createClient interfaces.ICreateClient
}

func NewClientsController(
	getClients interfaces.IGetClients,
	createClient interfaces.ICreateClient,
) *ClientsController {
	return &ClientsController{
		getClients:   getClients,
		createClient: createClient,
	}
}

// ListClients godoc
// @Summary      List clients
// @Description  get clients
// @Tags         clients
// @Accept       json
// @Produce      json
// @Success      200  {array}   responses.ClientsResponse
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      404  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /clients [get]
func (c *ClientsController) GetAll(ctx *gin.Context) {
	ucOutput, err := c.getClients.Perform()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses.GetClientsResponse(ucOutput))
}

// AddClient godoc
// @Summary      Add a client
// @Description  add by json client
// @Tags         clients
// @Accept       json
// @Produce      json
// @Param        client  body       payloads.ClientPayload  true  "Add client"
// @Success      200      {object}  responses.ClientsResponse
// @Failure      400      {object}  responses.ErrorResponse
// @Failure      404      {object}  responses.ErrorResponse
// @Failure      500      {object}  responses.ErrorResponse
// @Router       /clients [post]
func (c *ClientsController) Create(ctx *gin.Context) {
	var payload *payloads.ClientPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// colocar no arquivo de data
	parsedBirthDate, err := time.Parse(utils.DateFormat, payload.BirthDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientInput := &inputs.ClientInput{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		BirthDate: parsedBirthDate,
		Email:     payload.Email,
	}

	ucOutput, err := c.createClient.Perform(clientInput)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": utils.ToJson(ucOutput)})
}
