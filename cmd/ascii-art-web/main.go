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
	http.HandleFunc("/server.html", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving server.html from: docs/server.html")
		http.ServeFile(w, r, "docs/server.html")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Use /server.html or /ascii-art"))
	})
}

func main() {
	setupHandler()
	
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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

	bannerFile := "assets/" + req.Banner + ".txt"
	charMap, err := ascii.LoadBanner(bannerFile)
	if err != nil {
		sendError(w, "Banner not found", http.StatusNotFound)
		return
	}

	result := ascii.GenerateArtWithColorAndAlignment(req.Text, charMap, req.Substring, req.Color, req.Align)
	
	// Strip ANSI color codes for HTML display
	result = stripANSI(result)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Result: result})
}

func stripANSI(s string) string {
	result := ""
	inEscape := false
	for i := 0; i < len(s); i++ {
		if s[i] == '\033' && i+1 < len(s) && s[i+1] == '[' {
			inEscape = true
			i++
		} else if inEscape && s[i] == 'm' {
			inEscape = false
		} else if !inEscape {
			result += string(s[i])
		}
	}
	return result
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
