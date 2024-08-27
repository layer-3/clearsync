package quest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GinHandlePOST(c *gin.Context) {
	HandlePOST(c.Writer, c.Request)
}

func HandlePOST(w http.ResponseWriter, r *http.Request) {
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
		QuestID string `json:"quest_id"`
	}

	err = json.Unmarshal(body, &req)
	if err != nil || req.Address == "" || req.QuestID == "" {
		http.Error(w, "Invalid request or missing user address", http.StatusBadRequest)
		return
	}

	handler, exists := GetHandler(questKey, req.QuestID)
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

func GinHandleGET(c *gin.Context) {
	HandleGET(c.Writer, c.Request)
}

func HandleGET(w http.ResponseWriter, r *http.Request) {
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

	questID := r.URL.Query().Get("quest_id")
	if questID == "" {
		fmt.Println("error")
		http.Error(w, "quest_id is required", http.StatusBadRequest)
		return
	}

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
	// Input example: /galxe/balance
	segments := strings.Split(strings.Trim(path, "/"), "/")
	if len(segments) != 2 {
		return "", http.ErrNotSupported
	}
	// "galxe_balance"
	return segments[0] + "_" + segments[1], nil
}
