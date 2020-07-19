package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/broothie/battlesnake/robosnek/game"
	"google.golang.org/api/option"
)

var (
	start []byte
	games *firestore.CollectionRef

	logger = log.New(os.Stdout, "[battlesnake] ", log.LstdFlags)

	snakeConfig = map[string]string{
		"apiversion": "1",
		"author":     "broothie",
		"color":      "#ff00e6",
		"head":       "silly",
		"tail":       "bolt",
	}
)

func init() {
	var err error
	if start, err = json.Marshal(snakeConfig); err != nil {
		logger.Panicln(err)
	}

	ctx := context.Background()
	var options []option.ClientOption
	if os.Getenv("ENV") != "production" {
		options = append(options, option.WithCredentialsFile("battlesnake.gcloud-key.json"))
	}

	client, err := firestore.NewClient(ctx, "battlesnake-258923", options...)
	if err != nil {
		logger.Panicln(err)
	}

	games = client.Collection("games")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(start); err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/start", func(http.ResponseWriter, *http.Request) {})
	http.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var gameState *game.State
		if err := json.Unmarshal(data, &gameState); err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		gameState.Init(logger)
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprintf(w, `{"move":"%s"}`, gameState.NextMove()); err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/end", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var game map[string]interface{}
		if err := json.Unmarshal(data, &game); err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		game["created_at"] = time.Now()
		if _, _, err := games.Add(context.Background(), game); err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/games/last", func(w http.ResponseWriter, r *http.Request) {
		results, err := games.
			OrderBy("created_at", firestore.Desc).
			Limit(1).
			Documents(context.Background()).
			GetAll()
		if err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(results) < 1 {
			err := errors.New("no games found")
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		game := results[0].Data()
		id := game["game"].(map[string]interface{})["id"]
		data := map[string]interface{}{
			"id":         id,
			"iframe_url": fmt.Sprintf("https://board.battlesnake.com/?engine=https://engine.battlesnake.com&game=%s&autoplay=true&hideScoreboard=true&hideMediaControls=true&frameRate=6&loop=true", id),
		}

		logger.Println("iframe_url", data["iframe_url"])
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false)
		if err := enc.Encode(data); err != nil {
			logger.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Printf("serving @ %s\n", port)
	logger.Panicln(http.ListenAndServe(fmt.Sprintf(":%s", port), reqLogger(http.DefaultServeMux)))
}

func reqLogger(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		recorder := httptest.NewRecorder()
		before := time.Now()

		handler.ServeHTTP(recorder, r)

		since := time.Since(before)
		logger.Printf("%s %s %dB | %d %s %dB | %v \n", r.Method, r.URL.Path, r.ContentLength, recorder.Code, http.StatusText(recorder.Code), recorder.Body.Len(), since)

		for key, values := range recorder.Header() {
			for _, value := range values {
				w.Header().Add(key, value)
			}
		}

		w.WriteHeader(recorder.Code)
		if _, err := w.Write(recorder.Body.Bytes()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
