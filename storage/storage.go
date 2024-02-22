package storage

import (
	"noname_team_project/config"
	"noname_team_project/storage/neo4j"
	"noname_team_project/storage/elastic"
	"noname_team_project/storage/mongo"
	"noname_team_project/storage/postgre"
	"noname_team_project/storage/redis"
)

type Storage struct {
	Redis   *redis.Redis
	Postgre *postgre.Postgre
	Elastic *elastic.Elastic
	Neo4j   *neo4j.Neo4j
	Mongo   *mongo.Mongo
}

func New(config *config.Config) *Storage {
	return &Storage{
		Redis:   redis.New(config),
		Elastic: elastic.New(config),
		Postgre: postgre.New(config),
	}
}

func (s *Storage) Open() error {
	if err := s.Redis.Open(); err != nil {
		return err
	}
	if err := s.Postgre.Open(); err != nil {
		return err
	}
	s.Elastic.Open()
	s.Neo4j.Open()
	s.Mongo.Open()

	return nil
}
