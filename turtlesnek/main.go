package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var config = map[string]string{
	"apiversion": "1",
	"author":     "broothie",
	"color":      "#888888",
	"head":       "default",
	"tail":       "default",
}

func main() {
	bytes, err := json.Marshal(config)
	if err != nil {
		log.Panicln("failed to serialize snake config", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(bytes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/start", func(http.ResponseWriter, *http.Request) {})
	http.HandleFunc("/end", func(http.ResponseWriter, *http.Request) {})

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	log.Panicln(http.ListenAndServe(addr, nil))
}
