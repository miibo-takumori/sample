package service

import (
	"github.com/yourusername/calculator-api/internal/model"
	"github.com/yourusername/calculator-api/pkg/calculator"
)

// CalculatorService はビジネスロジックを提供します
type CalculatorService struct{}

// NewCalculatorService は新しいCalculatorServiceを作成します
func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

// Add は加算処理を実行します
func (s *CalculatorService) Add(req *model.CalculationRequest) (*model.CalculationResponse, error) {
	result, err := calculator.Add(req.Numbers)
	if err != nil {
		return nil, err
	}

	return &model.CalculationResponse{
		Result:    result,
		Operation: "addition",
	}, nil
}

// Subtract は減算処理を実行します
func (s *CalculatorService) Subtract(req *model.CalculationRequest) (*model.CalculationResponse, error) {
	result, err := calculator.Subtract(req.Numbers)
	if err != nil {
		return nil, err
	}

	return &model.CalculationResponse{
		Result:    result,
		Operation: "subtraction",
	}, nil
}

// Multiply は乗算処理を実行します
func (s *CalculatorService) Multiply(req *model.CalculationRequest) (*model.CalculationResponse, error) {
	result, err := calculator.Multiply(req.Numbers)
	if err != nil {
		return nil, err
	}

	return &model.CalculationResponse{
		Result:    result,
		Operation: "multiplication",
	}, nil
}

// Divide は除算処理を実行します
func (s *CalculatorService) Divide(req *model.CalculationRequest) (*model.CalculationResponse, error) {
	result, err := calculator.Divide(req.Numbers)
	if err != nil {
		return nil, err
	}

	return &model.CalculationResponse{
		Result:    result,
		Operation: "division",
	}, nil
}

// Average は平均値を計算します
func (s *CalculatorService) Average(req *model.CalculationRequest) (*model.CalculationResponse, error) {
	result, err := calculator.Average(req.Numbers)
	if err != nil {
		return nil, err
	}

	return &model.CalculationResponse{
		Result:    result,
		Operation: "average",
		Count:     len(req.Numbers),
	}, nil
}

// Median は中央値を計算します
func (s *CalculatorService) Median(req *model.CalculationRequest) (*model.CalculationResponse, error) {
	result, err := calculator.Median(req.Numbers)
	if err != nil {
		return nil, err
	}

	return &model.CalculationResponse{
		Result:    result,
		Operation: "median",
		Count:     len(req.Numbers),
	}, nil
}

// StandardDeviation は標準偏差を計算します
func (s *CalculatorService) StandardDeviation(req *model.CalculationRequest) (*model.CalculationResponse, error) {
	result, err := calculator.StandardDeviation(req.Numbers)
	if err != nil {
		return nil, err
	}

	return &model.CalculationResponse{
		Result:    result,
		Operation: "standard_deviation",
		Count:     len(req.Numbers),
	}, nil
}
