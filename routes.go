package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"io/ioutil"
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

	router.HandleFunc("/vg/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Upload File Hit")

		r.ParseMultipartForm(10 << 20)
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}).Methods(http.MethodPost)
}