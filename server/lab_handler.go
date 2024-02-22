package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type EsRequest struct {
	Phrase string
	Date   time.Time
}

func (s *Server) handleLab1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var esReq EsRequest

		err := json.NewDecoder(r.Body).Decode(&esReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		lectureList, err := s.storage.Elastic.GetByPhrase(esReq.Phrase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		studentArray, lessonsArray, err := s.storage.Neo4j.GetVisited(lectureList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		visitRate, err := s.storage.Postgre.GetVisitRate(studentArray, lessonsArray)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(visitRate)
	}
}
