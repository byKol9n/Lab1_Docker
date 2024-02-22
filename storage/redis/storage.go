package redis

import (
	"context"
	"fmt"
	"noname_team_project/config"
	rd "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type Redis struct {
	conf *config.Config
	conn *rd.Client
}

func New(config *config.Config) *Redis {
	return &Redis{
		conf: config,
	}
}

func (r *Redis) Open() error {
	conn := rd.NewClient(&rd.Options{
		Addr:     "localhost:6379",
		Password: r.conf.REDIS_PASSWORD,
		DB:       r.conf.REDIS_BD,
	})

	if err := conn.Ping(ctx).Err(); err != nil {
		fmt.Println("connect error")
		return err
	}

	r.conn = conn
	return nil
}

func (r *Redis) Close() {
	r.conn.Close()
}
