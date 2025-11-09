package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yourusername/calculator-api/internal/model"
	"github.com/yourusername/calculator-api/internal/service"
)

// CalculatorHandler はHTTPリクエストを処理します
type CalculatorHandler struct {
	service *service.CalculatorService
}

// NewCalculatorHandler は新しいCalculatorHandlerを作成します
func NewCalculatorHandler(service *service.CalculatorService) *CalculatorHandler {
	return &CalculatorHandler{
		service: service,
	}
}

// Add は加算のエンドポイント
func (h *CalculatorHandler) Add(w http.ResponseWriter, r *http.Request) {
	var req model.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Add(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// Subtract は減算のエンドポイント
func (h *CalculatorHandler) Subtract(w http.ResponseWriter, r *http.Request) {
	var req model.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Subtract(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// Multiply は乗算のエンドポイント
func (h *CalculatorHandler) Multiply(w http.ResponseWriter, r *http.Request) {
	var req model.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Multiply(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// Divide は除算のエンドポイント
func (h *CalculatorHandler) Divide(w http.ResponseWriter, r *http.Request) {
	var req model.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Divide(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// Average は平均値のエンドポイント
func (h *CalculatorHandler) Average(w http.ResponseWriter, r *http.Request) {
	var req model.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Average(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// Median は中央値のエンドポイント
func (h *CalculatorHandler) Median(w http.ResponseWriter, r *http.Request) {
	var req model.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.Median(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// StandardDeviation は標準偏差のエンドポイント
func (h *CalculatorHandler) StandardDeviation(w http.ResponseWriter, r *http.Request) {
	var req model.CalculationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	result, err := h.service.StandardDeviation(&req)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}

// respondWithJSON はJSONレスポンスを返します
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError はエラーレスポンスを返します
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, model.ErrorResponse{
		Error:   "error",
		Message: message,
	})
}
