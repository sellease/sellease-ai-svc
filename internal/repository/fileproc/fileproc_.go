package fileproc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sellease-ai/config"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/utils"
	"sellease-ai/logger"
)

func (r *fileProcRepository) AddProductListing(ctx context.Context, prodData models.ProductListing) (err error) {
	url := fmt.Sprintf("%s/products", config.GetConfig().SellEaseAPISvcUrl)

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(prodData)
	if err != nil {
		logger.WithContext(ctx).Errorf("Error marshaling JSON: %s", err.Error())
		return err
	}

	// Create a request with the POST method, URL, and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		logger.WithContext(ctx).Errorf("Error creating request: %s", err.Error())
		return err
	}

	// Set the Content-Type header to specify the JSON format
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := &http.Client{}

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		logger.WithContext(ctx).Errorf("Error sending request: %s", err.Error())
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.WithContext(ctx).Errorf("Error adding product listing")
		return utils.ErrAddingProductListing
	}

	return nil
}
