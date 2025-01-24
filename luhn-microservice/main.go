package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LuhnRequest struct {
	Number string `json:"number"`
}

type LuhnResponse struct {
	Valid bool `json:"valid"`
}

func luhnHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LuhnRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	valid := LuhnCheck(req.Number)
	resp := LuhnResponse{Valid: valid}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/check", luhnHandler)

	fmt.Println("Server listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
