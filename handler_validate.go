package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	tokens := strings.Fields(params.Body)
	banWords := []string{"kerfuffle", "sharbert", "fornax"}
	for idx := 0; idx < len(tokens); idx++ {
		if slices.Contains(banWords, strings.ToLower(tokens[idx])) {
			tokens[idx] = "****"
		}
	}

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: strings.Join(tokens, " "),
	})
}
