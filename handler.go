package main

import (
	"net/http"
)


func handlerTest(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{}{})
}

func handlerTestError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}