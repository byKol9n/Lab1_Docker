package elastic

import (
	"noname_team_project/config"
	"github.com/elastic/go-elasticsearch/v8"
)

type Elastic struct {
	conf *config.Config
	conn *elasticsearch.Client
}

func New(config *config.Config) *Elastic {
	return &Elastic{
		conf: config,
	}
}

func (e *Elastic) Open() error {
	cfg := elasticsearch.Config{
		Addresses: []string{
		  "http://localhost:9200",
		},
	  }
	  
	 conn, err := elasticsearch.NewClient(cfg)
	 if err != nil{
		return err
	 }

	 e.conn = conn
	 return nil
}

func (e *Elastic) Close() {
	e.conn = nil
}
