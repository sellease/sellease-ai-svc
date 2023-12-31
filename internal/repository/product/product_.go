package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sellease-ai/config"
	"sellease-ai/internal/entity/models"
	"sellease-ai/internal/entity/response"
	"sellease-ai/internal/utils"
	"sellease-ai/logger"
	"strings"
)

func (r *productRepository) GenerateProductDescription(ctx context.Context, data models.ProductDescriptionRequestData) (
	result response.ProductDescriptionResponse, err error) {

	url := config.GetConfig().TextCortexUrl
	apiKey := config.GetConfig().TextCortexAPIKey

	// Marshal the payload into JSON
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return result, err
	}

	// Create a request with the POST method, URL, and request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return result, err
	}

	// Set the Content-Type header to specify the JSON format
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Create an HTTP client
	client := &http.Client{}

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return result, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		fmt.Println("POST request was successful!")
	} else {
		fmt.Println("POST request failed. Status code:", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return result, err
	}

	// Parse the JSON response into the Response struct
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return result, err
	}

	return result, nil
}

func (r *productRepository) GenerateKeywords(ctx context.Context, value string) (result []string, err error) {
	// Define the URL
	url := config.GetConfig().RapidAPIKeywordUrl

	// Define query parameters
	country := "in"

	// Create a query string
	queryString := fmt.Sprintf("keyword=%s&country=%s", value, country)

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP GET request
	req, err := http.NewRequest("GET", url+"?"+queryString, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("X-RapidAPI-Key", config.GetConfig().RapidAPIKey)
	req.Header.Set("X-RapidAPI-Host", config.GetConfig().RapidAPIKeywordHost)

	// Send the GET request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the JSON response into the Response struct
	var response []models.KeywordData
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return result, err
	}

	for _, data := range response {
		result = append(result, data.Text)
	}

	return result, nil
}

func (r *productRepository) TranslateText(ctx context.Context, text, target string) (resultTxt string, err error) {
	// Define the URL
	reqUrl := config.GetConfig().GoogleTranslateUrl

	payload := url.Values{}
	payload.Add("q", text)
	payload.Add("target", target)
	payload.Add("source", "en")

	encodedPayload := payload.Encode()

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP POST request with the payload
	req, err := http.NewRequest("POST", reqUrl, strings.NewReader(encodedPayload))
	if err != nil {
		logger.WithContext(ctx).Errorf("Error creating request: %s", err.Error())
		return resultTxt, err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "application/gzip")
	req.Header.Set("X-RapidAPI-Key", config.GetConfig().RapidAPIKey)
	req.Header.Set("X-RapidAPI-Host", config.GetConfig().RapidAPITranslateHost)

	// Send the POST request
	resp, err := client.Do(req)
	if err != nil {
		logger.WithContext(ctx).Errorf("Error sending request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.WithContext(ctx).Errorf("Error in translate text endpoint [status_code: %d]", resp.StatusCode)
		return resultTxt, utils.ErrTranslation
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.WithContext(ctx).Errorf("Error reading response: %s", err.Error())
		return
	}

	var response models.GoogleTranslateRapidApiResp
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	return response.Data.Translations[0].TranslatedText, nil
}
