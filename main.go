package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/broothie/battlesnake/game"
)

var (
	start []byte

	logger = log.New(os.Stdout, "[battlesnake] ", log.LstdFlags)

	snakeConfig = map[string]string{
		"color":    "#ff00e6",
		"headType": "silly",
		"tailType": "bolt",
	}
)

func init() {
	var err error
	if start, err = json.Marshal(snakeConfig); err != nil {
		logger.Panicln(err)
	}
}

func main() {
	http.HandleFunc("/start", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(start)
	})

	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var gameState *game.State
		if err := json.Unmarshal(data, &gameState); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		gameState.Init(logger)

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"move":"%s"}`, gameState.NextMove())
	})

	http.HandleFunc("/end", func(http.ResponseWriter, *http.Request) {})
	http.HandleFunc("/ping", func(http.ResponseWriter, *http.Request) {})

	port := os.Getenv("PORT")
	logger.Printf("serving @ %s\n", port)
	logger.Panicln(http.ListenAndServe(fmt.Sprintf(":%s", port), reqLogger(http.DefaultServeMux)))
}

func reqLogger(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recorder := httptest.NewRecorder()
		before := time.Now()

		handler.ServeHTTP(recorder, r)

		since := time.Since(before)
		logger.Printf("%s %s | %d %s %d | %v \n", r.Method, r.URL.Path, recorder.Code, http.StatusText(recorder.Code), recorder.Body.Len(), since)

		for key, values := range recorder.Header() {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(recorder.Code)
		w.Write(recorder.Body.Bytes())
	}
}
