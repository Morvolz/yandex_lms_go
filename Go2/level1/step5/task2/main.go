package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Data string
	// тело ответа
	StatusCode int
	// код ответа
}

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	newCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(newCtx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResp := &APIResponse{
		Data:       string(bodyBytes),
		StatusCode: resp.StatusCode,
	}
	return apiResp, nil

}
