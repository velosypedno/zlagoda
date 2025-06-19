package utils

import (
	"testing"
)

func TestGenerateSecureID(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   bool // true if should succeed
	}{
		{
			name:   "valid length 10",
			length: 10,
			want:   true,
		},
		{
			name:   "valid length 13",
			length: 13,
			want:   true,
		},
		{
			name:   "valid length 1",
			length: 1,
			want:   true,
		},
		{
			name:   "valid length 50",
			length: 50,
			want:   true,
		},
		{
			name:   "zero length",
			length: 0,
			want:   true, // should work, just return empty string
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateSecureID(tt.length)

			if tt.want && err != nil {
				t.Errorf("GenerateSecureID() error = %v, wantErr %v", err, !tt.want)
				return
			}

			if len(got) != tt.length {
				t.Errorf("GenerateSecureID() length = %v, want %v", len(got), tt.length)
			}

			// Check that generated ID contains only valid characters
			if tt.length > 0 {
				for _, char := range got {
					if !isValidIDChar(char) {
						t.Errorf("GenerateSecureID() contains invalid character: %c", char)
					}
				}
			}
		})
	}
}

func TestGenerateSecureID_Uniqueness(t *testing.T) {
	const iterations = 1000
	const idLength = 10

	generated := make(map[string]bool)

	for i := 0; i < iterations; i++ {
		id, err := GenerateSecureID(idLength)
		if err != nil {
			t.Fatalf("GenerateSecureID() failed at iteration %d: %v", i, err)
		}

		if generated[id] {
			t.Errorf("GenerateSecureID() generated duplicate ID: %s", id)
		}

		generated[id] = true
	}
}

func TestIsDecimalValid(t *testing.T) {
	tests := []struct {
		name                   string
		value                  float64
		maxDigitsBeforeDecimal int
		maxDigitsAfterDecimal  int
		want                   bool
	}{
		{
			name:                   "valid small number",
			value:                  123.45,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  2,
			want:                   true,
		},
		{
			name:                   "valid whole number",
			value:                  12345,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  2,
			want:                   true,
		},
		{
			name:                   "valid decimal only",
			value:                  0.99,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  2,
			want:                   true,
		},
		{
			name:                   "too many digits before decimal",
			value:                  123456,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  2,
			want:                   false,
		},
		{
			name:                   "too many digits after decimal",
			value:                  123.456,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  2,
			want:                   false,
		},
		{
			name:                   "negative number",
			value:                  -123.45,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  2,
			want:                   false,
		},
		{
			name:                   "zero",
			value:                  0,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  2,
			want:                   true,
		},
		{
			name:                   "very small positive number",
			value:                  0.0001,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  4,
			want:                   true,
		},
		{
			name:                   "edge case - exactly max digits",
			value:                  99999.9999,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  4,
			want:                   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDecimalValid(tt.value, tt.maxDigitsBeforeDecimal, tt.maxDigitsAfterDecimal)
			if got != tt.want {
				t.Errorf("IsDecimalValid(%v, %d, %d) = %v, want %v",
					tt.value, tt.maxDigitsBeforeDecimal, tt.maxDigitsAfterDecimal, got, tt.want)
			}
		})
	}
}

func TestIsSalaryValid(t *testing.T) {
	tests := []struct {
		name   string
		salary float64
		want   bool
	}{
		{
			name:   "valid salary",
			salary: 50000.00,
			want:   true,
		},
		{
			name:   "valid salary with decimals",
			salary: 1234.56,
			want:   true,
		},
		{
			name:   "maximum valid salary",
			salary: 9999999999.9999,
			want:   true,
		},
		{
			name:   "zero salary",
			salary: 0,
			want:   true,
		},
		{
			name:   "negative salary",
			salary: -1000,
			want:   false,
		},
		{
			name:   "too many digits before decimal",
			salary: 12345678901.00, // 11 digits before decimal
			want:   false,
		},
		{
			name:   "too many decimal places",
			salary: 1000.12345, // 5 decimal places
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSalaryValid(tt.salary)
			if got != tt.want {
				t.Errorf("IsSalaryValid(%v) = %v, want %v", tt.salary, got, tt.want)
			}
		})
	}
}

func TestIsAmountValid(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
		want   bool
	}{
		{
			name:   "valid amount",
			amount: 199.99,
			want:   true,
		},
		{
			name:   "valid large amount",
			amount: 1000000,
			want:   true,
		},
		{
			name:   "zero amount",
			amount: 0,
			want:   true,
		},
		{
			name:   "valid small amount",
			amount: 0.01,
			want:   true,
		},
		{
			name:   "negative amount",
			amount: -50.00,
			want:   false,
		},
		{
			name:   "too many digits",
			amount: 12345678901.00, // 11 digits before decimal
			want:   false,
		},
		{
			name:   "too precise",
			amount: 100.12345, // 5 decimal places
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAmountValid(tt.amount)
			if got != tt.want {
				t.Errorf("IsAmountValid(%v) = %v, want %v", tt.amount, got, tt.want)
			}
		})
	}
}

// Benchmark tests
func BenchmarkGenerateSecureID_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GenerateSecureID(10)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerateSecureID_13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GenerateSecureID(13)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkIsDecimalValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsDecimalValid(12345.6789, 10, 4)
	}
}

func BenchmarkIsSalaryValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsSalaryValid(50000.00)
	}
}

func BenchmarkIsAmountValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAmountValid(199.99)
	}
}

// Helper function to check if character is valid for ID
func isValidIDChar(char rune) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9')
}

// Test edge cases and error conditions
func TestGenerateSecureID_EdgeCases(t *testing.T) {
	// Test with very large length
	id, err := GenerateSecureID(1000)
	if err != nil {
		t.Errorf("GenerateSecureID(1000) failed: %v", err)
	}
	if len(id) != 1000 {
		t.Errorf("GenerateSecureID(1000) length mismatch: got %d, want 1000", len(id))
	}
}

func TestIsDecimalValid_EdgeCases(t *testing.T) {
	tests := []struct {
		name                   string
		value                  float64
		maxDigitsBeforeDecimal int
		maxDigitsAfterDecimal  int
		want                   bool
	}{
		{
			name:                   "zero max digits before",
			value:                  0.5,
			maxDigitsBeforeDecimal: 0,
			maxDigitsAfterDecimal:  2,
			want:                   false, // 0.5 has 1 digit before decimal (0)
		},
		{
			name:                   "zero max digits after",
			value:                  123.0,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  0,
			want:                   true, // 123.0 is effectively 123
		},
		{
			name:                   "trailing zeros after decimal",
			value:                  123.1000,
			maxDigitsBeforeDecimal: 5,
			maxDigitsAfterDecimal:  1,
			want:                   true, // Should trim trailing zeros
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsDecimalValid(tt.value, tt.maxDigitsBeforeDecimal, tt.maxDigitsAfterDecimal)
			if got != tt.want {
				t.Errorf("IsDecimalValid(%v, %d, %d) = %v, want %v",
					tt.value, tt.maxDigitsBeforeDecimal, tt.maxDigitsAfterDecimal, got, tt.want)
			}
		})
	}
}
