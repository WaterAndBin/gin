package utils

import (
	"encoding/json"
)

func TransJson(v any) string {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "json.MarshalIndent error: " + err.Error()
	}
	return string(data)
}
