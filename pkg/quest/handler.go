package quest

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandlePOST(c *gin.Context) {
	w := c.Writer
	r := c.Request

	questKey, err := extractQuestKey(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid path format", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	var req struct {
		Address string `json:"address"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil || req.Address == "" {
		http.Error(w, "Invalid request or missing user address", http.StatusBadRequest)
		return
	}

	questID := c.Param("id")
	handler, exists := GetHandler(questKey, questID)
	if !exists {
		http.Error(w, "Handler not found for given quest ID", http.StatusNotFound)
		return
	}

	result, err := handler.Handle(r.Context(), req.Address)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func HandleGET(c *gin.Context) {
	w := c.Writer
	r := c.Request

	questKey, err := extractQuestKey(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid path format", http.StatusBadRequest)
		return
	}

	userAddress := r.URL.Query().Get("address")
	if userAddress == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}

	questID := c.Param("id")
	handler, exists := GetHandler(questKey, questID)
	if !exists {
		http.Error(w, "Handler not found for given quest ID", http.StatusNotFound)
		return
	}

	result, err := handler.Handle(r.Context(), userAddress)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func extractQuestKey(path string) (string, error) {
	// Input example: /galxe/balance/1
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) != 3 {
		return "", http.ErrNotSupported
	}
	// "galxe_balance"
	return segments[0] + "_" + segments[1], nil
}
