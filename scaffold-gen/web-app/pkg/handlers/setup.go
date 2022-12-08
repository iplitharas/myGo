package handlers

import "net/http"

func CreateMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", healthCheck)
	mux.HandleFunc("/api", api)
	mux.HandleFunc("/", index)
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
	return mux
}
