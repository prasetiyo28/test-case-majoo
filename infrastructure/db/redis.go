package db

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewDatabase(address string, password string, db string) (*RedisCache, error) {
	dbVal, _ := strconv.Atoi(db)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       dbVal,
	})
	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return &RedisCache{
		Client: client,
	}, nil
}
