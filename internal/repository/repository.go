package repository

import (
	"sellease-ai/internal/repository/product"
)

type Repository struct {
	// User    user.RepositoryInterface
	Product product.RepositoryInterface
}

func InitRepository() *Repository {
	return &Repository{
		// User:    user.InitUserRepository(db),
		Product: product.InitProductRepository(),
	}
}
