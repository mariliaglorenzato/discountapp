package usecases

import (
	domain "discountapp/domain"
	"discountapp/usecases/interfaces"
)

type GetProducts struct {
	Repository interfaces.IRepository
}

func NewGetProducts(repository interfaces.IRepository) interfaces.IGetProducts {
	return &GetProducts{Repository: repository}
}

func (usecase *GetProducts) Perform() ([]*domain.Product, error) {
	output, err := usecase.Repository.GetAllProducts()
	if err != nil {
		return nil, err
	}

	return output, nil
	// 	{
	// 		ID: 1, FullName: "John Coltrane",
	// 		// ArtWorks: []domain.ArtWork{
	// 		// 	{ID: 1, Name: "Blue Train", Description: "Test.....", PublishedAt: time.Now()},
	// 		// },
	// 	},
	// 	{
	// 		ID: 2, FullName: "Gerry Mulligan",
	// 		// ArtWorks: []domain.ArtWork{
	// 		// 	{ID: 2, Name: "Jeru", Description: "Test.....", PublishedAt: time.Now()},
	// 		// },
	// 	},
	// 	{
	// 		ID: 3, FullName: "Sarah Vaughan",
	// 		// ArtWorks: []domain.ArtWork{
	// 		// 	{ID: 3, Name: "Sarah Vaughan and Clifford Brown", Description: "Test...", PublishedAt: time.Now()},
	// 		// },
	// 	},
	// }
}
