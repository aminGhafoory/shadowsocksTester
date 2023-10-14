package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aminghafoory/shadowTester/internal/database"
)

func (apicfg *apiConfig) addSub(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := Sub{}
	decoder.Decode(&params)
	sub, err := apicfg.DB.AddToSubs(context.Background(), database.AddToSubsParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       params.URL,
	})
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error: can not add sub %+v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, dbSubToSub(sub))
}

func (apiCfg *apiConfig) showSubs(w http.ResponseWriter, r *http.Request) {
	dbsubs, err := apiCfg.DB.GetAllSubs(context.Background())
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	subs := []Sub{}
	for _, dbsub := range dbsubs {
		subs = append(subs, dbSubToSub(dbsub))
	}
	respondWithJSON(w, http.StatusOK, subs)
}
