package model

// CalculationRequest は計算リクエストの構造体
type CalculationRequest struct {
	Numbers []float64 `json:"numbers"`
}

// CalculationResponse は計算結果のレスポンス構造体
type CalculationResponse struct {
	Result    float64 `json:"result"`
	Operation string  `json:"operation"`
	Count     int     `json:"count,omitempty"`
}

// ErrorResponse はエラーレスポンスの構造体
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
