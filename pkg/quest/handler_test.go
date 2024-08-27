package quest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock handler for testing
type MockHandler struct{}

func (h *MockHandler) Handle(ctx context.Context, userAddress string) (bool, error) {
	return true, nil
}

func init() {
	// Register mock handler for testing
	RegisterHandler("galxe_balance", "1", &MockHandler{})
}

func TestPostHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		body           io.Reader
		expectedStatus int
		expectedBody   bool
	}{
		{
			name:           "Valid POST Request",
			method:         http.MethodPost,
			url:            "/galxe/balance",
			body:           bytes.NewBuffer([]byte(`{"address": "0xUserAddress", "quest_id": "1"}`)),
			expectedStatus: http.StatusOK,
			expectedBody:   true,
		},
		{
			name:           "Invalid Path Format",
			method:         http.MethodPost,
			url:            "/galxe",
			body:           bytes.NewBuffer([]byte(`{"address": "0xUserAddress", "quest_id": "1"}`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Missing Body",
			method:         http.MethodPost,
			url:            "/galxe/balance",
			body:           nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Invalid Body",
			method:         http.MethodPost,
			url:            "/galxe/balance",
			body:           bytes.NewBuffer([]byte(`{}`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Handler Not Found",
			method:         http.MethodPost,
			url:            "/galxe/unknown",
			body:           bytes.NewBuffer([]byte(`{"address": "0xUserAddress", "quest_id": "1"}`)),
			expectedStatus: http.StatusNotFound,
			expectedBody:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, tt.body)
			rec := httptest.NewRecorder()

			HandlePOST(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Decode JSON response and compare with expected body
			var actualBody bool
			json.NewDecoder(rec.Body).Decode(&actualBody)

			if tt.expectedStatus == http.StatusOK && actualBody != tt.expectedBody {
				t.Errorf("expected valid %v, got %v", tt.expectedBody, actualBody)
			}
		})
	}
}

func TestGetHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		expectedStatus int
		expectedBody   bool
	}{
		{
			name:           "Valid GET Request",
			method:         http.MethodGet,
			url:            "/galxe/balance?address=0xUserAddress&quest_id=1",
			expectedStatus: http.StatusOK,
			expectedBody:   true,
		},
		{
			name:           "Invalid Path Format",
			method:         http.MethodGet,
			url:            "/galxe",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Missing User Address",
			method:         http.MethodGet,
			url:            "/galxe/balance",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Handler Not Found",
			method:         http.MethodGet,
			url:            "/galxe/balance?address=0xUserAddress&quest_id=2",
			expectedStatus: http.StatusNotFound,
			expectedBody:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, nil)
			rec := httptest.NewRecorder()

			HandleGET(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			// Decode JSON response and compare with expected body
			var actualBody bool
			json.NewDecoder(rec.Body).Decode(&actualBody)

			if tt.expectedStatus == http.StatusOK && actualBody != tt.expectedBody {
				t.Errorf("expected valid %v, got %v", tt.expectedBody, actualBody)
			}
		})
	}
}
