package persistence

import (
	"discountapp/domain"
	"discountapp/usecases/interfaces"

	"gorm.io/gorm"
)

/*
* Melhorias: separar em ProductRepository e ClientRepository
 */
type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) interfaces.IRepository {
	return &Repository{db: db}
}

func (r *Repository) GetAllProducts() ([]*domain.Product, error) {
	products := []*domain.Product{}
	result := r.db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *Repository) GetProduct(product *domain.Product) (*domain.Product, error) {
	result := r.db.Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *Repository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	result := r.db.Create(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *Repository) GetClient(client *domain.Client) (*domain.Client, error) {
	result := r.db.Find(&client)

	if result.Error != nil {
		return nil, result.Error
	}

	return client, nil
}

func (r *Repository) GetAllClients() ([]*domain.Client, error) {
	clients := []*domain.Client{}
	result := r.db.Find(&clients)

	if result.Error != nil {
		return nil, result.Error
	}

	return clients, nil
}

func (r *Repository) CreateClient(client *domain.Client) (*domain.Client, error) {
	result := r.db.Create(&client)
	if result.Error != nil {
		return nil, result.Error
	}

	return client, nil
}
