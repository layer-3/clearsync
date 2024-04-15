package quotes

import (
	"encoding/json"
	"io"
	"net/http"
)

func getMapping(url string) (map[string][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var mappings map[string]map[string][]string
	if err := json.Unmarshal(body, &mappings); err != nil {
		return nil, err
	}
	return mappings["tokens"], nil
}
