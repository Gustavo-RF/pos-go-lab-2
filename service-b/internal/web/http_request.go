package web

import (
	"context"
	"io"
	"net/http"
)

func Request(url, method string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func RequestWithContext(ctx context.Context, url, method string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Close = true

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return res, nil
}
