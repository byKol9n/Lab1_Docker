package main

import (
	"log"
	"noname_team_project/server"
)

func main() {
	s := server.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}