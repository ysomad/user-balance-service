package v1

import (
	"encoding/json"
	"net/http"
)

func writeError(w http.ResponseWriter, code int, err error, details map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(Error{
		Status:  http.StatusText(code),
		Message: err.Error(),
		Details: details,
	})
}

func writeOK(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(v)
}
