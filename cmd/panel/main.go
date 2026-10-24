package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type healthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Time    string `json:"time"`
}

func main() {
	addr := os.Getenv("PANEL_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(healthResponse{
			Status:  "ok",
			Service: "panel-conecta-web-ssh",
			Time:    time.Now().UTC().Format(time.RFC3339),
		})
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		_, _ = w.Write([]byte(`<!doctype html>
<html lang="es">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width,initial-scale=1">
<title>PANEL CONECTA WEB SSH</title>
<style>
body{margin:0;background:#0b0b12;color:#eee;font-family:system-ui,-apple-system,sans-serif;display:grid;min-height:100vh;place-items:center}
main{text-align:center;padding:32px}
h1{margin:0 0 10px}
p{color:#a9a9b8}
.badge{display:inline-block;padding:7px 12px;border-radius:999px;background:#171725;color:#b9a7ff}
</style>
</head>
<body><main><span class="badge">ONLINE</span><h1>PANEL CONECTA WEB SSH</h1><p>Base del nuevo panel web.</p></main></body>
</html>`))
	})

	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Printf("PANEL CONECTA WEB SSH listening on %s", addr)
	log.Fatal(server.ListenAndServe())
}
