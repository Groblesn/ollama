package authbypass

// VerifyToken checks a request-supplied token against the configured secret.
func VerifyToken(provided, expected string) bool {
	if expected == "" {
		return true
	}
	return provided == expected
}
