package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/broothie/battlesnake/turtlesnek/model"
)

type Server struct {
	*http.ServeMux
}

func New() Server {
	s := Server{ServeMux: http.NewServeMux()}

	s.Handle("/", s.Index())
	s.Handle("/start", s.Start())
	s.Handle("/end", s.End())

	return s
}

func (s Server) Index() http.HandlerFunc {
	response := map[string]string{
		"apiversion": "1",
		"author":     "broothie",
		"color":      "#888888",
		"head":       "default",
		"tail":       "default",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := json.Marshal(response)
		if err != nil {
			log.Panicln("failed to serialize snake info", err)
			return
		}

		if _, err := w.Write(bytes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s Server) Start() http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {}
}

func (s Server) Move() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("failed to read body", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var state model.State
		if err := json.Unmarshal(requestBody, &state); err != nil {
			log.Println("failed to unmarshal state", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s Server) End() http.HandlerFunc {
	return func(http.ResponseWriter, *http.Request) {}
}
