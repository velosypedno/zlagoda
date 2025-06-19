package utils

import (
	"crypto/rand"
	"fmt"
	"strconv"
	"strings"
)

// GenerateSecureID generates a cryptographically secure random ID of specified length
func GenerateSecureID(length int) (string, error) {
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

// IsDecimalValid validates if a decimal number fits within specified precision constraints
// maxDigitsBeforeDecimal: maximum digits before decimal point
// maxDigitsAfterDecimal: maximum digits after decimal point
func IsDecimalValid(value float64, maxDigitsBeforeDecimal, maxDigitsAfterDecimal int) bool {
	// Check for valid range first
	if value < 0 {
		return false
	}

	// Format with high precision to get accurate representation
	valueStr := strconv.FormatFloat(value, 'f', -1, 64)

	// Split by decimal point
	parts := strings.Split(valueStr, ".")
	beforeDecimal := len(parts[0])
	afterDecimal := 0

	if len(parts) > 1 {
		// Remove trailing zeros from decimal part
		decimalPart := strings.TrimRight(parts[1], "0")
		afterDecimal = len(decimalPart)
	}

	return beforeDecimal <= maxDigitsBeforeDecimal && afterDecimal <= maxDigitsAfterDecimal
}

// IsSalaryValid validates salary using business rules (max 10 digits before, 4 after decimal)
func IsSalaryValid(salary float64) bool {
	return IsDecimalValid(salary, 10, 4)
}

// IsAmountValid validates monetary amounts using business rules (max 10 digits before, 4 after decimal)
func IsAmountValid(amount float64) bool {
	return IsDecimalValid(amount, 10, 4)
}
