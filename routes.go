package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutes(router *mux.Router) {
	// enableCORS(router)

	router.HandleFunc("/vg", func(w http.ResponseWriter, r *http.Request){
		videoGames, err := all()
		if err == nil {
			respondWithSuccess(videoGames, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("vg/{id}", func(w http.ResponseWriter, r *http.Request) {
		_id := mux.Vars(r)["id"]
		id, err := stringToInt64(_id)

		if err != nil {
			respondWithError(err, w)
			return
		}

		videoGame, err := get(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(videoGame, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/vg", func(w http.ResponseWriter, r *http.Request) {
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)
		
		if err != nil {
			respondWithError(err, w)
		} else {
			err := create(videoGame)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/vg", func(w http.ResponseWriter, r *http.Request) {
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)

		if err != nil {
			respondWithError(err, w)
		} else {
			err := update(videoGame)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)

	router.HandleFunc("/vg/{id}", func(w http.ResponseWriter, r *http.Request) {
		_id := mux.Vars(r)["id"]
		id, err := stringToInt64(_id)

		if err != nil {
			respondWithError(err, w)
			return
		}

		err = delete(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}