package hmactrunc

import (
	"crypto/hmac"
	"crypto/sha256"
)

func Verify(key, message, providedTag []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expected := mac.Sum(nil)
	if len(providedTag) < 8 {
		return false
	}
	return hmac.Equal(expected[:8], providedTag[:8])
}
