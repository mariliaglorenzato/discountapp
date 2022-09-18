package usecases_test

import (
	"errors"

	mocks "discountapp/tests/mocks/usecases/interfaces"

	"discountapp/domain"
	"discountapp/tests/factories"
	"discountapp/usecases"
	"discountapp/usecases/inputs"
	"discountapp/usecases/outputs"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Getting Discount Data", func() {
	var getEntityProduct func() *domain.Product
	var getEntityClient func() *domain.Client
	var getInput func() inputs.DiscountInput
	var getProductMockedError func() error
	var getClientMockedError func() error
	var expectedOutput *outputs.DiscountOutput
	var output *outputs.DiscountOutput
	var err error
	var product *domain.Product
	var client *domain.Client

	BeforeEach(func() {
		getInput = func() inputs.DiscountInput {
			return inputs.DiscountInput{
				ProductSlug: "caneta-de-cd",
				ClientEmail: "test@gmail.com",
			}
		}

		getEntityProduct = func() *domain.Product {
			if product == nil {
				product = factories.BuildProduct()
			}
			return product
		}
		getEntityClient = func() *domain.Client {
			if client == nil {
				client = factories.BuildClient()
			}
			return client
		}
		getProductMockedError = func() error {
			return nil
		}
		getClientMockedError = func() error {
			return nil
		}
	})
	JustBeforeEach(func() {
		product := getEntityProduct()
		client := getEntityClient()
		input := getInput()
		productMockedError := getProductMockedError()
		clientMockedError := getClientMockedError()
		repositoryMock := &mocks.IRepository{}
		usecase := usecases.NewGetDiscount(repositoryMock)
		repositoryMock.On("GetProduct", &domain.Product{Slug: input.ProductSlug}).Return(product, productMockedError)
		repositoryMock.On("GetClient", &domain.Client{Email: input.ClientEmail}).Return(client, clientMockedError)
		output, err = usecase.Perform(&input)
	})

	When("use case retrieves discount", func() {
		When("It successfully retrieve discount", func() {
			BeforeEach(func() {
				expectedOutput = &outputs.DiscountOutput{
					ProductTile:          "Caneta de CD",
					ProductPrice:         2874, // 0,042 * 3000
					OriginalProductPrice: 3000,
					DiscountPercentage:   4.2,
				}
			})

			It("Returns expected output", func() {
				Expect(output).To(Equal(expectedOutput))
				Expect(err).To(BeNil())
			})
		})

		When("Does not retrieve discount", func() {
			When("It does not find client", func() {
				BeforeEach(func() {
					getClientMockedError = func() error {
						return errors.New("Not Found")
					}
				})

				It("Returns error", func() {
					Expect(err).To(Equal(getClientMockedError()))
					Expect(output).To(BeNil())
				})
			})
			When("It does not find the product", func() {
				BeforeEach(func() {
					getProductMockedError = func() error {
						return errors.New("")
					}
				})

				It("Returns expected output", func() {
					Expect(err).To(Equal(getProductMockedError()))
					Expect(output).To(BeNil())
				})
			})
		})
	})
})
