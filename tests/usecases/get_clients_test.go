package usecases_test

import (
	"discountapp/domain"
	"discountapp/tests/factories"
	mocks "discountapp/tests/mocks/usecases/interfaces"
	"discountapp/usecases"
	"discountapp/usecases/outputs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lists all clients", func() {
	var getClientsList func() []*domain.Client
	var getMockedError func() error
	var output *outputs.ClientsOutput
	var clientsList []*domain.Client
	var err error

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
	})
	JustBeforeEach(func() {
		clientsList := getClientsList()
		// stubbedProduct := factories.BuildStubbedProduct()
		// input := getInput()
		mockedError := getMockedError()
		repositoryMock := &mocks.IRepository{}
		usecase := usecases.NewGetClients(repositoryMock)
		repositoryMock.On("GetAllClients").Return(clientsList, mockedError)
		output, err = usecase.Perform()
	})

	It("gets all clients", func() {
		Expect(output).To(Equal(&outputs.ClientsOutput{
			Clients: getClientsList(),
		}))
	})

	When("there is no client", func() {
		BeforeEach(func() {
			getClientsList = func() []*domain.Client {
				return []*domain.Client{}
			}
		})

		It("Returns expected error", func() {
			Expect(output).To(Equal(&outputs.ClientsOutput{
				Clients: []*domain.Client{},
			}))
			Expect(err).To(BeNil())
		})
	})
})
