package utils

const (
	// Base62 characters for URL encoding
	base62Chars     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortCodeLength = 7
)

func GenerateShortCode() (string, error) {
	// TODO: Generate random base62 encoded short code
	// TODO: Use cryptographically secure random number generator
	return "", nil
}

func GenerateCustomShortCode(customCode string) (string, error) {
	// TODO: Validate and sanitize custom short code
	// TODO: Check length and allowed characters
	return "", nil
}

func EncodeBase62(num int64) string {
	// TODO: Encode number to base62 string
	return ""
}

func DecodeBase62(str string) (int64, error) {
	// TODO: Decode base62 string to number
	return 0, nil
}

func GenerateAPIKey() (string, error) {
	// TODO: Generate secure API key
	// TODO: Use 32-byte random string encoded as hex
	return "", nil
}

func GenerateRandomString(length int) (string, error) {
	// TODO: Generate random string of specified length
	// TODO: Use base62 character set
	return "", nil
}

func ValidateShortCode(code string) bool {
	// TODO: Validate short code format
	// TODO: Check length and character set
	return false
}
