package quest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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

	var req struct {
		Address common.Address `json:"address" binding:"required"`
	}

	if err := c.BindJSON(&req); err != nil {
		http.Error(w, "Invalid request or missing user address", http.StatusBadRequest)
		return
	}

	questID := c.Param("id")
	handler, exists := GetHandler(questKey, questID)
	if !exists {
		http.Error(w, "Handler not found for given quest ID", http.StatusNotFound)
		return
	}

	result, err := handler.Handle(r.Context(), req.Address.Hex())
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

	var req struct {
		Address string `form:"address" binding:"required"`
	}

	if err := c.Bind(&req); err != nil {
		http.Error(w, "Invalid request or missing user address", http.StatusBadRequest)
		return
	}

	userAddress := common.HexToAddress(req.Address)

	questID := c.Param("id")
	handler, exists := GetHandler(questKey, questID)
	if !exists {
		http.Error(w, "Handler not found for given quest ID", http.StatusNotFound)
		return
	}

	result, err := handler.Handle(r.Context(), userAddress.Hex())
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
