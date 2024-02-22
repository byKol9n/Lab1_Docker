package mongo

import "noname_team_project/config"

type Mongo struct {
	conf *config.Config
}

func New(config *config.Config) *Mongo {
	return &Mongo{
		conf: config,
	}
}

func (e *Mongo) Open() error {
	return nil
}

func (e *Mongo) Close() {
}
