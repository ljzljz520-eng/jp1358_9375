package api

import (
	"net/http/httptest"
	"path/filepath"
	"testing"
	"yoga.example/studio/internal/service"
	"yoga.example/studio/internal/storage"
)

func TestHealth(t *testing.T) {
	st, _ := storage.Open(filepath.Join(t.TempDir(), "db"))
	defer st.Close()
	r := httptest.NewRecorder()
	New(service.New(st)).Handler().ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
