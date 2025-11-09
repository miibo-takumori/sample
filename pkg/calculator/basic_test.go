package calculator

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
		wantErr  bool
	}{
		{"Simple addition", []float64{1, 2, 3}, 6, false},
		{"Negative numbers", []float64{-5, 10, -3}, 2, false},
		{"Single number", []float64{42}, 42, false},
		{"Empty array", []float64{}, 0, true},
		{"Decimals", []float64{1.5, 2.5, 3.0}, 7.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Add(tt.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Add() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
		wantErr  bool
	}{
		{"Simple subtraction", []float64{10, 3, 2}, 5, false},
		{"Negative result", []float64{5, 10}, -5, false},
		{"Single number", []float64{42}, 42, false},
		{"Empty array", []float64{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Subtract(tt.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Subtract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Subtract() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
		wantErr  bool
	}{
		{"Simple multiplication", []float64{2, 3, 4}, 24, false},
		{"With zero", []float64{5, 0, 10}, 0, false},
		{"Negative numbers", []float64{-2, 3, -4}, 24, false},
		{"Empty array", []float64{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Multiply(tt.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Multiply() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []float64
		expected float64
		wantErr  bool
	}{
		{"Simple division", []float64{100, 2, 5}, 10, false},
		{"Division by zero", []float64{10, 0}, 0, true},
		{"Single number", []float64{42}, 42, false},
		{"Empty array", []float64{}, 0, true},
		{"Decimal result", []float64{10, 4}, 2.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.numbers)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Divide() = %v, want %v", result, tt.expected)
			}
		})
	}
}
