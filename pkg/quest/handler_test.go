package quest

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
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
			url:            "/galxe/balance/1",
			body:           bytes.NewBuffer([]byte(`{"address": "0x0a61E07824C1D2b7596C4E787280930E36444509"}`)),
			expectedStatus: http.StatusOK,
			expectedBody:   true,
		},
		{
			name:           "Invalid POST Request data",
			method:         http.MethodPost,
			url:            "/galxe/balance/1",
			body:           bytes.NewBuffer([]byte(`{"address": "0xQa61e07824c1d2b7596c4e787280930e36444509"}`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   true,
		},
		{
			name:           "Invalid Path Format",
			method:         http.MethodPost,
			url:            "/galxe",
			body:           bytes.NewBuffer([]byte(`{"address": "0x0a61E07824C1D2b7596C4E787280930E36444509"}`)),
			expectedStatus: http.StatusNotFound,
			expectedBody:   false,
		},
		{
			name:           "Missing Body",
			method:         http.MethodPost,
			url:            "/galxe/balance/1",
			body:           nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Invalid Body",
			method:         http.MethodPost,
			url:            "/galxe/balance/1",
			body:           bytes.NewBuffer([]byte(`{}`)),
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Handler Not Found",
			method:         http.MethodPost,
			url:            "/galxe/unknown",
			body:           bytes.NewBuffer([]byte(`{"address": "0x0a61E07824C1D2b7596C4E787280930E36444509"}`)),
			expectedStatus: http.StatusNotFound,
			expectedBody:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			router.POST("/galxe/balance/:id", HandlePOST)

			req := httptest.NewRequest(tt.method, tt.url, tt.body)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, rec.Code)
			}

			if tt.expectedStatus == http.StatusOK {
				var response bool
				err := json.NewDecoder(rec.Body).Decode(&response)
				require.NoError(t, err)

				if response != tt.expectedBody {
					t.Errorf("expected valid %v, got %v", tt.expectedBody, response)
				}
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
			url:            "/galxe/balance/1?address=0x0a61E07824C1D2b7596C4E787280930E36444509",
			expectedStatus: http.StatusOK,
			expectedBody:   true,
		},
		{
			name:           "Invalid Path Format",
			method:         http.MethodGet,
			url:            "/galxe",
			expectedStatus: http.StatusNotFound,
			expectedBody:   false,
		},
		{
			name:           "Missing User Address",
			method:         http.MethodGet,
			url:            "/galxe/balance/1",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   false,
		},
		{
			name:           "Handler Not Found",
			method:         http.MethodGet,
			url:            "/galxe/balance/2?address=0x0a61E07824C1D2b7596C4E787280930E36444509",
			expectedStatus: http.StatusNotFound,
			expectedBody:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize a new Gin router
			router := gin.Default()
			router.GET("/galxe/balance/:id", HandleGET)

			req := httptest.NewRequest(tt.method, tt.url, nil)
			rec := httptest.NewRecorder()

			// Serve the HTTP request using the router
			router.ServeHTTP(rec, req)

			// Assert the expected status code
			require.Equal(t, tt.expectedStatus, rec.Code)

			// Decode JSON response if the status is OK
			if tt.expectedStatus == http.StatusOK {
				var response bool
				err := json.NewDecoder(rec.Body).Decode(&response)
				require.NoError(t, err)

				// Check the response body
				require.Equal(t, tt.expectedBody, response)
			}
		})
	}
}
