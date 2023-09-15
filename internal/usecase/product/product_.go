package product

import (
	"context"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/entity/response"
	"sellease-ai/logger"
)

func (u *productUsecase) GenerateProductDesc(ctx context.Context, req request.ProductDescriptionRequest) (
	result response.Output, err error) {

	data := models.ProductDescriptionRequestData{
		Brand:       req.Brand,
		Category:    req.Category,
		Description: req.Description,
		Keywords:    req.Keywords,
		MaxTokens:   512,
		Model:       "chat-sophos-1",
		N:           1,
		Name:        req.Name,
		SourceLang:  "en",
		TargetLang:  "en",
		Temperature: 0.65,
	}

	resp, err := u.productRepo.GenerateProductDescription(ctx, data)
	if err != nil {
		logger.Errorf("error while generating product description - %s", err.Error())
		return result, err
	}

	logger.Infof("remaining credits - %f", resp.Data.RemainingCredits)

	for _, output := range resp.Data.Outputs {
		result.Index = output.Index
		result.Text = output.Text
		result.ID = output.ID
	}

	return result, err
}
