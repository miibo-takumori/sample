package calculator

import (
	"errors"
	"math"
	"sort"
)

// Average は数値の平均を計算します
func Average(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	sum, _ := Add(numbers)
	return sum / float64(len(numbers)), nil
}

// Median は数値の中央値を計算します
func Median(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	// コピーしてソート
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	n := len(sorted)
	if n%2 == 0 {
		// 偶数個の場合は中央2つの平均
		return (sorted[n/2-1] + sorted[n/2]) / 2, nil
	}
	// 奇数個の場合は中央の値
	return sorted[n/2], nil
}

// StandardDeviation は数値の標準偏差を計算します
func StandardDeviation(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	avg, _ := Average(numbers)

	var sumSquaredDiff float64
	for _, num := range numbers {
		diff := num - avg
		sumSquaredDiff += diff * diff
	}

	variance := sumSquaredDiff / float64(len(numbers))
	return math.Sqrt(variance), nil
}

// Min は最小値を返します
func Min(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	min := numbers[0]
	for _, num := range numbers[1:] {
		if num < min {
			min = num
		}
	}
	return min, nil
}

// Max は最大値を返します
func Max(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New("numbers array is empty")
	}

	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}
	return max, nil
}
