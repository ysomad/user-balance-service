package v1

import (
	"encoding/json"
	"net/http"
)

func writeError(w http.ResponseWriter, code int, err error, details map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Error{
		Code:    http.StatusText(code),
		Message: err.Error(),
		Details: details,
	})
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
	w.WriteHeader(code)
}
