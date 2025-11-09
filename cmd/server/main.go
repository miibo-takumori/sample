package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yourusername/calculator-api/internal/handler"
	"github.com/yourusername/calculator-api/internal/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// サービス層の初期化
	calcService := service.NewCalculatorService()

	// ハンドラー層の初期化
	calcHandler := handler.NewCalculatorHandler(calcService)

	// ルーターの設定
	r := mux.NewRouter()

	// ヘルスチェック
	r.HandleFunc("/health", healthCheckHandler).Methods("GET")

	// API v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// 基本計算
	api.HandleFunc("/calculate/add", calcHandler.Add).Methods("POST")
	api.HandleFunc("/calculate/subtract", calcHandler.Subtract).Methods("POST")
	api.HandleFunc("/calculate/multiply", calcHandler.Multiply).Methods("POST")
	api.HandleFunc("/calculate/divide", calcHandler.Divide).Methods("POST")

	// 統計計算
	api.HandleFunc("/statistics/average", calcHandler.Average).Methods("POST")
	api.HandleFunc("/statistics/median", calcHandler.Median).Methods("POST")
	api.HandleFunc("/statistics/stddev", calcHandler.StandardDeviation).Methods("POST")

	// サーバー起動
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
