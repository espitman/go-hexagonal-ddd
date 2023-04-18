package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Connection struct {
	Addr     string
	Password string
	DB       int
}

func NewConnection(addr string, password string, db int) *Connection {
	return &Connection{
		Addr:     addr,
		Password: password,
		DB:       db,
	}

}

func (c *Connection) NewClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return client, nil
}
