package fileproc

import (
	"context"
	"sellease-ai/internal/entity/models"
)

type RepositoryInterface interface {
	AddProductListing(ctx context.Context, prodData models.ProductListing) (err error)
}

type fileProcRepository struct {
}

func InitfileProcRepository() RepositoryInterface {
	return &fileProcRepository{}
}
