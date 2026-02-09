package main

import (
	"ascii-art/internal/ascii"
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Text      string `json:"text"`
	Banner    string `json:"banner"`
	Color     string `json:"color,omitempty"`
	Substring string `json:"substring,omitempty"`
	Align     string `json:"align,omitempty"`
}

type Response struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func setupHandler() {
	http.HandleFunc("/ascii-art", asciiArtHandler)
}

func main() {
	setupHandler()
	
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendError(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		sendError(w, "Text is required", http.StatusBadRequest)
		return
	}

	if req.Banner == "" {
		req.Banner = "standard"
	}

	if req.Align != "" && !isValidAlignment(req.Align) {
		sendError(w, "Invalid alignment", http.StatusBadRequest)
		return
	}

	bannerFile := "../../assets/" + req.Banner + ".txt"
	charMap, err := ascii.LoadBanner(bannerFile)
	if err != nil {
		sendError(w, "Banner not found", http.StatusNotFound)
		return
	}

	result := ascii.GenerateArtWithColorAndAlignment(req.Text, charMap, req.Substring, req.Color, req.Align)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: result})
}

func isValidAlignment(align string) bool {
	validAlignments := []string{"left", "right", "center", "justify"}
	for _, valid := range validAlignments {
		if align == valid {
			return true
		}
	}
	return false
}

func sendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
