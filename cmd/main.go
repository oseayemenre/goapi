package main

import (
	"log"
	"net/http"
)

type api struct {
	addr string
}

func (a *api) getHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home Route"))
}

func main() {
	api:=&api{addr: "8000"}
	
	mux := http.NewServeMux()

	svr := http.Server{
		Addr: api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /home", api.getHandler)

	if err := svr.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}