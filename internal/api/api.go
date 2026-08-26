package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"yoga.example/studio/internal/model"
	"yoga.example/studio/internal/service"
)

type Server struct {
	Svc *service.Service
	Mux *http.ServeMux
}

func New(s *service.Service) *Server {
	a := &Server{Svc: s, Mux: http.NewServeMux()}
	a.routes()
	return a
}
func (a *Server) routes() {
	a.Mux.HandleFunc("/health", a.health)
	a.Mux.HandleFunc("/records", a.records)
	a.Mux.HandleFunc("/profile", a.profile)
	a.Mux.HandleFunc("/upload/", a.upload)
}
func (a *Server) Handler() http.Handler { return logging(a.Mux) }
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func (a *Server) health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
func (a *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var x model.Record
		if json.NewDecoder(r.Body).Decode(&x) != nil {
			http.Error(w, "bad json", 400)
			return
		}
		if e := a.Svc.Register(x); e != nil {
			http.Error(w, e.Error(), 422)
			return
		}
		json.NewEncoder(w).Encode(x)
		return
	}
	id := r.URL.Query().Get("id")
	x, e := a.Svc.Store.GetRecord(id)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(x)
}
func (a *Server) profile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "method", 405)
		return
	}
	var p model.Profile
	if json.NewDecoder(r.Body).Decode(&p) != nil {
		http.Error(w, "bad json", 400)
		return
	}
	if e := a.Svc.Profile(p); e != nil {
		http.Error(w, e.Error(), 422)
		return
	}
	json.NewEncoder(w).Encode(p)
}
func (a *Server) upload(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/upload/")
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	res, e := a.Svc.UploadCover(ctx, id, []byte("cover"))
	if e != nil {
		http.Error(w, res.Message(), 499)
		return
	}
	json.NewEncoder(w).Encode(res)
}
