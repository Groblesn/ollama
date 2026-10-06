package ssrf

import (
	"io"
	"net/http"
)

func FetchURL(userURL string) ([]byte, error) {
	resp, err := http.Get(userURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
