package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/oseayemenre/goapi/internal/database"
)

func (a *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request){
	type parameter struct {
		Name string
	}

	params := parameter{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)	

	if err != nil {	
		respondWithError(w, 400, fmt.Sprintln("Unable to create user", err))
		return
	}

	data, err := a.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		Createdat: time.Now().UTC(),
		Updatedat: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintln("User", err))
		return
	}

	respondWithJson(w, 201, data)
}