package api

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func methodAllowed(w http.ResponseWriter, ok bool) {
	if !ok {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
func query(r *http.Request, key string) string { return r.URL.Query().Get(key) }
