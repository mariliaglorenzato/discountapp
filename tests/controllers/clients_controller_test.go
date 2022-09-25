package controllers_test

import (
	"net/http/httptest"

	"discountapp/controllers"
	"discountapp/controllers/responses"
	"discountapp/domain"
	"discountapp/tests/factories"
	mocks "discountapp/tests/mocks/usecases/interfaces"
	"discountapp/tests/utils"
	"discountapp/usecases/outputs"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Clients controller", func() {
	// var output *outputs.ClientsOutput
	var getClientsList func() []*domain.Client
	var getMockedError func() error
	var clientsList []*domain.Client
	var getResponse func() *httptest.ResponseRecorder
	var responseRecorder *httptest.ResponseRecorder

	BeforeEach(func() {
		getClientsList = func() []*domain.Client {
			if clientsList == nil {
				clientsList = factories.BuildClients()
			}
			return clientsList
		}
		getMockedError = func() error {
			return nil
		}
		getResponse = func() *httptest.ResponseRecorder {
			if responseRecorder == nil {
				responseRecorder = httptest.NewRecorder()
			}
			return responseRecorder
		}
	})
	JustBeforeEach(func() {
		responseRecorder := getResponse()
		ctx, _ := gin.CreateTestContext(responseRecorder)
		clientsList := getClientsList()
		mockedError := getMockedError()
		// repositoryMock := &mocks.IRepository{}
		mockedGetClients := &mocks.IGetClients{}
		mockedCreateClient := &mocks.ICreateClient{}

		clientsOutput := &outputs.ClientsOutput{
			Clients: clientsList,
		}

		mockedGetClients.On("Perform").Return(clientsOutput, mockedError)

		clientsController := controllers.NewClientsController(mockedGetClients, mockedCreateClient)
		clientsController.GetAll(ctx)
	})

	When("lists all clients", func() {
		It("returns 200 status code with body response", func() {
			data := utils.LoadFile[responses.ClientsResponse]("../fixtures/get_clients.json")

			Expect(200).To(Equal(responseRecorder.Code))
			Expect(utils.ToJson(data)).To(Equal(responseRecorder.Body.String()))
		})

		When("there is no client", func() {
			BeforeEach(func() {
			})

			It("Returns expected error", func() {
			})
		})
	})
})
