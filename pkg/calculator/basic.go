package calculator

import "errors"

// Add は数値の合計を計算します
func Add(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum, nil
}

// Subtract は最初の数値から残りの数値を順に引きます
func Subtract(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result -= numbers[i]
	}
	return result, nil
}

// Multiply は数値の積を計算します
func Multiply(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result *= numbers[i]
	}
	return result, nil
}

// Divide は最初の数値を残りの数値で順に割ります
func Divide(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] == 0 {
			return 0, errors.New("division by zero")
		}
		result /= numbers[i]
	}
	return result, nil
}
