package repository

import (
	"database/sql"

	"github.com/Aibekabdi/rest-api/internal/models"
	"github.com/go-redis/redis/v8"
)

type Counter interface {
	Add(num string) error
	Sub(num string) error
	Val() (int, error)
}
type User interface {
	PostUserindb(user models.UserModel) (int, error)
}
type Repository struct {
	Counter
	User
}

func NewRepository(client *redis.Client, db *sql.DB) *Repository {
	return &Repository{
		Counter: NewCounterRedis(client),
		User:    NewUserSqlite(db),
	}
}
