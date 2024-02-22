package postgre

import (
	"database/sql"
	"fmt"
	"noname_team_project/config"

	_ "github.com/lib/pq"
)

type Postgre struct {
	conf *config.Config
	conn *sql.DB
}

func New(config *config.Config) *Postgre {
	return &Postgre{
		conf: config,
	}
}

func (p *Postgre) Open() error {
	conn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", p.conf.POSTGRES_USER, p.conf.POSTGRES_PASSWORD, "postgre", "5432", p.conf.POSTGRES_DB))
	if err != nil {
		return err
	}

	if err = conn.Ping(); err != nil {
		return err
	}

	p.conn = conn
	return nil
}

func (p *Postgre) Close() {
	p.conn.Close()
}
