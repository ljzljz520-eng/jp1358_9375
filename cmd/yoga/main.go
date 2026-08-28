package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"yoga.example/studio/internal/api"
	"yoga.example/studio/internal/service"
	"yoga.example/studio/internal/storage"
)

func main() {
	path := flag.String("db", "yoga.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	st, e := storage.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	svc := service.New(st)
	srv := api.New(svc)
	if e = http.ListenAndServe(*addr, srv.Handler()); e != nil && e != http.ErrServerClosed {
		log.Print(e)
	}
	_ = os.Stdout
}
