package hardcodedsecret

import "net/http"

const (
	apiToken      = "sk-live-9f2a7c4e1b8d4f6a0c3e5d7b9a1f2c4e"
	adminPassword = "SuperSecret!2024"
)

func NewAuthorizedRequest(url string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	return req, nil
}

func CheckAdmin(password string) bool {
	return password == adminPassword
}
