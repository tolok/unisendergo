package unisendergo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func request(ctx context.Context, client *http.Client, url string, data interface{}, response interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(j)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, reader)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		var responseError ErrorResponse

		if err := json.Unmarshal(body, &responseError); err != nil {
			return err
		}

		return errors.New(responseError.Message)
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	return nil
}
