package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAsciiArtHandler_Success(t *testing.T) {
	req := Request{Text: "Hi", Banner: "standard"}
	body, _ := json.Marshal(req)
	
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestAsciiArtHandler_DefaultBanner(t *testing.T) {
	req := Request{Text: "Hi"}
	body, _ := json.Marshal(req)
	
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestAsciiArtHandler_WithColor(t *testing.T) {
	req := Request{Text: "Hi", Banner: "shadow", Color: "red"}
	body, _ := json.Marshal(req)
	
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestAsciiArtHandler_WithSubstring(t *testing.T) {
	req := Request{Text: "Hello", Banner: "thinkertoy", Color: "blue", Substring: "ell"}
	body, _ := json.Marshal(req)
	
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestAsciiArtHandler_WithAlignment(t *testing.T) {
	alignments := []string{"left", "right", "center", "justify"}
	for _, align := range alignments {
		req := Request{Text: "Hi", Banner: "standard", Align: align}
		body, _ := json.Marshal(req)
		
		r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
		w := httptest.NewRecorder()
		
		asciiArtHandler(w, r)
		
		if w.Code != http.StatusOK {
			t.Errorf("Expected 200 for align=%s, got %d", align, w.Code)
		}
	}
}

func TestAsciiArtHandler_MethodNotAllowed(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestAsciiArtHandler_InvalidJSON(t *testing.T) {
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader([]byte("invalid")))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestAsciiArtHandler_EmptyText(t *testing.T) {
	req := Request{Text: "", Banner: "standard"}
	body, _ := json.Marshal(req)
	
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestAsciiArtHandler_InvalidAlignment(t *testing.T) {
	req := Request{Text: "Hi", Banner: "standard", Align: "invalid"}
	body, _ := json.Marshal(req)
	
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
}

func TestAsciiArtHandler_BannerNotFound(t *testing.T) {
	req := Request{Text: "Hi", Banner: "nonexistent"}
	body, _ := json.Marshal(req)
	
	r := httptest.NewRequest(http.MethodPost, "/ascii-art", bytes.NewReader(body))
	w := httptest.NewRecorder()
	
	asciiArtHandler(w, r)
	
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected 404, got %d", w.Code)
	}
}

func TestIsValidAlignment(t *testing.T) {
	tests := []struct {
		align string
		valid bool
	}{
		{"left", true},
		{"right", true},
		{"center", true},
		{"justify", true},
		{"invalid", false},
		{"", false},
	}
	
	for _, tt := range tests {
		result := isValidAlignment(tt.align)
		if result != tt.valid {
			t.Errorf("isValidAlignment(%q) = %v, want %v", tt.align, result, tt.valid)
		}
	}
}

func TestSendError(t *testing.T) {
	w := httptest.NewRecorder()
	sendError(w, "test error", http.StatusBadRequest)
	
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected 400, got %d", w.Code)
	}
	
	var errResp ErrorResponse
	json.NewDecoder(w.Body).Decode(&errResp)
	
	if errResp.Error != "test error" {
		t.Errorf("Expected 'test error', got %q", errResp.Error)
	}
}

func TestSetupHandler(t *testing.T) {
	setupHandler()
}
