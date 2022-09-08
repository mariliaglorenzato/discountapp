package interfaces

import "discountapp/domain"

type IRepository interface {
	GetAllProducts() ([]*domain.Product, error)
	GetAllClients() ([]*domain.Client, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	CreateClient(client *domain.Client) (*domain.Client, error)
	GetClient(client *domain.Client) (*domain.Client, error)
	GetProduct(product *domain.Product) (*domain.Product, error)
}
