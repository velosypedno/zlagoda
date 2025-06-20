package utils

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
)

func GenerateUPC(length int) (string, error) {
	const charset = "0123456789"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b), nil
}

func GenerateID(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	for i := range b {
		b[i] = charset[b[i]%byte(len(charset))]
	}
	return string(b), nil
}

func IsDecimalValid(value float64) bool {
	if value < 0 {
		return false
	}

	var valueStr string = strconv.FormatFloat(value, 'f', -1, 64)

	var parts []string = strings.Split(valueStr, ".")
	var beforeDecimal int = len(parts[0])
	var afterDecimal int = 0

	if len(parts) > 1 {
		decimalPart := strings.TrimRight(parts[1], "0")
		afterDecimal = len(decimalPart)
	}

	return beforeDecimal <= 13 && afterDecimal <= 4
}
