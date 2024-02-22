package server

import (
	"io"
	"net/http"
	"noname_team_project/config"
	"noname_team_project/storage"

	"github.com/gorilla/mux"
)

type Server struct {
	router  *mux.Router
	storage *storage.Storage
}

func New() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureStore(); err != nil {
		return err
	}
	s.configureRouter()

	return http.ListenAndServe(":4040", s.router)
}

func (s *Server) configureStore() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}

	st := storage.New(cfg)
	if err := st.Open(); err != nil {
		return err
	}

	s.storage = st
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleIndex())
	s.router.HandleFunc("/lab1", s.handleLab1())
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "server response")
	}
}
