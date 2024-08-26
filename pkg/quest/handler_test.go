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

func (h *MockHandler) Handle(ctx context.Context, userAddress string) (Result, error) {
	return Result{
		Valid: true,
		Data:  map[string]interface{}{"token_balance": 10.0},
	}, nil
}

func init() {
	// Register mock handler for testing
	RegisterHandler("galxe_balance", &MockHandler{})
}

func TestPostHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		body           io.Reader
		expectedStatus int
		expectedBody   Result
	}{
		{
			name:           "Valid POST Request",
			method:         http.MethodPost,
			url:            "/galxe/balance",
			body:           bytes.NewBuffer([]byte(`{"address": "0xUserAddress"}`)),
			expectedStatus: http.StatusOK,
			expectedBody: Result{
				Valid: true,
				Data:  map[string]interface{}{"token_balance": 10.0},
			},
		},
		{
			name:           "Invalid Path Format",
			method:         http.MethodPost,
			url:            "/galxe",
			body:           bytes.NewBuffer([]byte(`{"address": "0xUserAddress"}`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Result{},
		},
		{
			name:           "Missing Body",
			method:         http.MethodPost,
			url:            "/galxe/balance",
			body:           nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Result{},
		},
		{
			name:           "Invalid Body",
			method:         http.MethodPost,
			url:            "/galxe/balance",
			body:           bytes.NewBuffer([]byte(`{}`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Result{},
		},
		{
			name:           "Handler Not Found",
			method:         http.MethodPost,
			url:            "/galxe/unknown",
			body:           bytes.NewBuffer([]byte(`{"address": "0xUserAddress"}`)),
			expectedStatus: http.StatusNotFound,
			expectedBody:   Result{},
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
			var actualBody Result
			json.NewDecoder(rec.Body).Decode(&actualBody)

			if tt.expectedStatus == http.StatusOK && actualBody.Valid != tt.expectedBody.Valid {
				t.Errorf("expected valid %v, got %v", tt.expectedBody.Valid, actualBody.Valid)
			}

			if tt.expectedStatus == http.StatusOK && actualBody.Data["token_balance"] != tt.expectedBody.Data["token_balance"] {
				t.Errorf("expected token_balance %v, got %v", tt.expectedBody.Data["token_balance"], actualBody.Data["token_balance"])
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
		expectedBody   Result
	}{
		{
			name:           "Valid GET Request",
			method:         http.MethodGet,
			url:            "/galxe/balance?address=0xUserAddress",
			expectedStatus: http.StatusOK,
			expectedBody: Result{
				Valid: true,
				Data:  map[string]interface{}{"token_balance": 10.0},
			},
		},
		{
			name:           "Invalid Path Format",
			method:         http.MethodGet,
			url:            "/galxe",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Result{},
		},
		{
			name:           "Missing User Address",
			method:         http.MethodGet,
			url:            "/galxe/balance",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   Result{},
		},
		{
			name:           "Handler Not Found",
			method:         http.MethodGet,
			url:            "/galxe/unknown?address=0xUserAddress",
			expectedStatus: http.StatusNotFound,
			expectedBody:   Result{},
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
			var actualBody Result
			json.NewDecoder(rec.Body).Decode(&actualBody)

			if tt.expectedStatus == http.StatusOK && actualBody.Valid != tt.expectedBody.Valid {
				t.Errorf("expected valid %v, got %v", tt.expectedBody.Valid, actualBody.Valid)
			}

			if tt.expectedStatus == http.StatusOK && actualBody.Data["token_balance"] != tt.expectedBody.Data["token_balance"] {
				t.Errorf("expected token_balance %v, got %v", tt.expectedBody.Data["token_balance"], actualBody.Data["token_balance"])
			}
		})
	}
}
