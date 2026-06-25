package main

import "testing"

func TestLuhnValid(t *testing.T) {
	tests := []struct {
		name     string
		card     string
		expected bool
	}{
		{"Valid Visa", "4539 1488 0343 6467", true},
		{"Valid Mastercard", "5555 5555 5555 4444", true},
		{"Valid Amex", "3782 8224 6310 005", true},
		{"Valid Discover", "6011 1111 1111 1117", true},
		{"Valid Diners", "3056 9309 0259 04", true},
		{"Invalid", "1234 5678 9012 3456", false},
		{"Too short", "5", false},
		{"Empty", "", false},
		{"With letters", "4539a1488b0343c6467", true},
		{"Valid test card", "4111 1111 1111 1111", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := LuhnValid(tt.card)
			if res != tt.expected {
				t.Errorf("LuhnValid(%s) = %v, want %v", tt.card, res, tt.expected)
			}

			res = LuhnValidLUT(tt.card)
			if res != tt.expected {
				t.Errorf("LuhnValidLUT(%s) = %v, want %v", tt.card, res, tt.expected)
			}
		})
	}
}
