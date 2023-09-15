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
	var emptyString string

	data := models.ProductDescriptionRequestData{
		Brand:       emptyString,
		Category:    emptyString,
		Description: emptyString,
		Keywords:    []string{req.Name},
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

func (u *productUsecase) GenerateKeywords(ctx context.Context, value string) (result []string, err error) {
	result, err = u.productRepo.GenerateKeywords(ctx, value)
	if err != nil {
		logger.Errorf("error while generating keywords - %s", err.Error())
	}

	return result, err
}

func (u *productUsecase) TranslateText(ctx context.Context, request request.TranslationRequest) (
	res response.TranslateTextResponse, err error) {

	res.Name, err = u.productRepo.TranslateText(ctx, request.Name, request.Language)
	if err != nil {
		logger.Errorf("error while translating text - %s", err.Error())
		return res, err
	}
	res.Description, err = u.productRepo.TranslateText(ctx, request.Description, request.Language)
	if err != nil {
		logger.Errorf("error while translating text - %s", err.Error())
		return res, err
	}
	res.Ingredients, err = u.productRepo.TranslateText(ctx, request.Ingredients, request.Language)
	if err != nil {
		logger.Errorf("error while translating text - %s", err.Error())
		return res, err
	}
	res.HowToUse, err = u.productRepo.TranslateText(ctx, request.HowToUse, request.Language)
	if err != nil {
		logger.Errorf("error while translating text - %s", err.Error())
		return res, err
	}
	return
}
