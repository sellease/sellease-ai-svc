package product

import (
	"context"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/entity/response"
	"sellease-ai/internal/repository/product"
)

type UsecaseInterface interface {
	GenerateProductDesc(ctx context.Context, req request.ProductDescriptionRequest) (
		result response.Output, err error)
	GenerateKeywords(ctx context.Context, value string) (result []string, err error)
}

type productUsecase struct {
	productRepo product.RepositoryInterface
}

func InitProductUsecase(
	product product.RepositoryInterface,
) UsecaseInterface {
	return &productUsecase{
		productRepo: product,
	}
}
