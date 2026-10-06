package ssrfallow

import (
	"errors"
	"io"
	"net/http"
	"strings"
)

func Fetch(rawURL string) ([]byte, error) {
	if !strings.Contains(rawURL, "trusted.example.com") {
		return nil, errors.New("host not allowed")
	}
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
