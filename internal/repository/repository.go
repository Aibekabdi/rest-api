package repository

import (
	"github.com/go-redis/redis/v8"
)

type Counter interface {
	Add(num string) error
	Sub(num string) error
	Val() (int, error)
}
type Repository struct {
	Counter
}

func NewRepository(client *redis.Client) *Repository {
	return &Repository{Counter: NewCounterRedis(client)}
}
