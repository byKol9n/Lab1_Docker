package neo4j

import (
	"context"
	"noname_team_project/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4j struct {
	conf    *config.Config
	context context.Context
	conn    neo4j.DriverWithContext
}

func New(config *config.Config) *Neo4j {
	return &Neo4j{
		conf:    config,
		context: context.Background(),
	}
}

func (n *Neo4j) Open() error {
	driver, err := neo4j.NewDriverWithContext("bolt://localhost", neo4j.BasicAuth(n.conf.NEO4J_USER, n.conf.NEO4J_PASSWORD, ""))
	err = driver.VerifyConnectivity(n.context)
	if err != nil {
		panic(err)
	}

	defer driver.Close(n.context)

	n.conn = driver
	return nil
}

func (n *Neo4j) Close() {
}
