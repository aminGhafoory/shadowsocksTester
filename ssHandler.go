package main

import (
	"context"
	"net/http"
)

func (apiCfg *apiConfig) BestConfigs(w http.ResponseWriter, r *http.Request) {
	dbbestList, err := apiCfg.DB.GetBestList(context.Background())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	bestlist := []Bestlink{}
	for _, v := range dbbestList {
		bestlist = append(bestlist, dbBestLinkToBestlink(v))
	}
	respondWithJSON(w, 200, bestlist)
}

func (apiCfg *apiConfig) ShowAllSSs(w http.ResponseWriter, r *http.Request) {
	dbSSs, err := apiCfg.DB.GetAllSSs(context.Background())
	if err != nil {
		respondWithError(w, 400, err.Error())
	}
	SSs := []Shadowsocks{}
	for _, dbSS := range dbSSs {
		SSs = append(SSs, DbSStoShadowsocks(dbSS))
	}
	respondWithJSON(w, 200, SSs)
}
