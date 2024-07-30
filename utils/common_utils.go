package utils

import (
	"io"
	"net/http"
)

func MakeCurlCall(request_type, url string) ([]byte, error) {

	var resp *http.Response
	var err error

	switch request_type {
	case "GET":
		resp, err = http.Get(url)
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
