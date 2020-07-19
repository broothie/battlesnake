package server

import (
	"fmt"
	"net/http"
)

type MoveFunc func(jsonBytes []byte) (move []byte)

type Server struct {
	*http.ServeMux
	SnakeConfig []byte
	MoveFunc    MoveFunc
}

func New(snakeConfig []byte, moveFunc MoveFunc) *Server {
	server := &Server{
		ServeMux:    http.NewServeMux(),
		SnakeConfig: snakeConfig,
		MoveFunc:    moveFunc,
	}

	server.Handle("/", method(http.MethodGet, server.Index))
	return server
}

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write(s.SnakeConfig); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func method(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, fmt.Sprintf("invalid method for this route '%s'", r.Method), http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	}
}
