package fileproc

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"sellease-ai/consts"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/entity/request"
	"sellease-ai/internal/entity/response"
	"sellease-ai/internal/utils"
	"sellease-ai/logger"
	"strconv"
)

func (uc *fileProcUsecase) ProcessFile(ctx context.Context, req request.FileUploadRequest) (
	resp response.ProcessFileResponse, err error) {

	file, err := req.CSVFile.Open()
	if err != nil {
		logger.WithContext(ctx).Errorf("unable to open uploaded file - %s", err.Error())
		return resp, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()
	if err != nil {
		logger.WithContext(ctx).Errorf("error reading file - %s", err.Error())
		return resp, err
	}

	if l := len(rows); l < 1 {
		logger.WithContext(ctx).Errorf("empty file - %s", err.Error())
		return resp, err
	} else {
		if err := validateFileHeaders(consts.GetFileProcHeaders(), rows[0]); err != nil {
			logger.WithContext(ctx).Errorf("invalid file - %s", err.Error())
			return resp, err
		}
		for i := 1; i < len(rows); i++ {
			parseRes, err := parseRecord(rows[i])
			if err != nil {
				resp.ErrorCount++
				resp.Errors = append(resp.Errors, parseRes.(response.FileError))
				continue
			}
			prodListing := parseRes.(models.ProductListing)
			err = uc.fileProcRepo.AddProductListing(ctx, prodListing)
			if err != nil {
				resp.ErrorCount++
				continue
			}
			resp.SuccessCount++
		}
	}

	return resp, nil
}

func validateFileHeaders(requiredHeaders, actualHeaders []string) error {
	for _, required := range requiredHeaders {
		isFound := false
		for _, actual := range actualHeaders {
			if required == actual {
				isFound = true
				break
			}
		}
		if !isFound {
			return fmt.Errorf("column %s not found", required)
		}
	}
	return nil
}

func parseRecord(row []string) (any, error) {
	record := models.ProductListing{}

	for index, value := range row {
		switch index {
		case 0:
			record.SKU = value
		case 1:
			record.Name = value
		case 2:
			stockLimit, err := strconv.Atoi(value)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			record.StockLimit = stockLimit
		case 3:
			version, err := strconv.Atoi(value)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			record.Version = version
		case 4:
			retailPrice, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			record.RetailPrice = retailPrice
		case 5:
			record.Description = value
			if len(record.Description) < 100 {
				return response.FileError{
					SKU:   record.SKU,
					Error: utils.ErrDescriptionTooShort.Error(),
				}, utils.ErrDescriptionTooShort
			}
		case 6:
			record.HowToUse = value
		case 7:
			record.Ingredients = value
		case 8:
			// productImagesJson := fmt.Sprintf("`%s`", value)
			err := json.Unmarshal([]byte(value), &record.ProductImages)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			if len(record.ProductImages) < 4 {
				return response.FileError{
					SKU:   record.SKU,
					Error: utils.ErrProductImagesInsufficient.Error(),
				}, utils.ErrProductImagesInsufficient
			}
		case 9:
			height, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			record.Height = height
		case 10:
			width, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			record.Width = width
		case 11:
			length, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			record.Length = length
		case 12:
			weight, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return response.FileError{
					SKU:   record.SKU,
					Error: err.Error(),
				}, err
			}
			record.Weight = weight
		}
	}
	record.BrandID = "676394f1-6a62-4042-93a5-76b70ffab1d2"
	record.CategoryID = "24e8b3d3-92c7-40ca-ac5c-86451c61f846"
	return record, nil
}
