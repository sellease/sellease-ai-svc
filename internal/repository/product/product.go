package product

import (
	"context"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/entity/response"
)

type RepositoryInterface interface {
	GenerateProductDescription(ctx context.Context, data models.ProductDescriptionRequestData) (
		result response.ProductDescriptionResponse, err error)
	GenerateKeywords(ctx context.Context, value string) (result []string, err error)
}

type productRepository struct {
}

func InitProductRepository() RepositoryInterface {
	return &productRepository{}
}
