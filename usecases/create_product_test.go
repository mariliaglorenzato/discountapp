package usecases_test

import (
	"errors"

	"discountapp/domain"
	mocks "discountapp/mocks/usecases/interfaces"
	"discountapp/tests/factories"
	"discountapp/usecases"
	"discountapp/usecases/inputs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Creating a new Product", func() {
	var getEntityProduct func() *domain.Product
	var getInput func() inputs.ProductInput
	var getMockedError func() error
	var output *domain.Product
	var err error
	var product *domain.Product

	BeforeEach(func() {
		getInput = func() inputs.ProductInput {
			return inputs.ProductInput{
				Title:       "Caneta de CD",
				Description: "Caneta preta não sai com alcool",
				Price:       30,
			}
		}

		getEntityProduct = func() *domain.Product {
			if product == nil {
				product = factories.BuildProduct()
			}
			return product
		}
		getMockedError = func() error {
			return nil
		}
	})
	JustBeforeEach(func() {
		product := getEntityProduct()
		stubbedProduct := factories.BuildStubbedProduct()
		input := getInput()
		mockedError := getMockedError()
		repositoryMock := &mocks.IRepository{}
		usecase := usecases.NewCreateProduct(repositoryMock)
		repositoryMock.On("CreateProduct", stubbedProduct).Return(product, mockedError)
		output, err = usecase.Perform(&input)
	})

	It("creates a new product", func() {
		Expect(output).To(Equal(getEntityProduct()))
	})

	When("use case fails on create product", func() {
		BeforeEach(func() {
			getMockedError = func() error {
				return errors.New("Error while creating new product")
			}
		})

		It("Returns expected error", func() {
			Expect(output).To(BeNil())
			Expect(err).To(Equal(getMockedError()))
		})
	})
})
