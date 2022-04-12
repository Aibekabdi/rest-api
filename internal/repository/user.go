package repository

import (
	"database/sql"

	"github.com/Aibekabdi/rest-api/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type UserSqlite struct {
	db *sql.DB
}

func NewUserSqlite(db *sql.DB) *UserSqlite {
	return &UserSqlite{db: db}
}

func (u *UserSqlite) PostUserindb(user models.UserModel) (int, error) {
	query, err := u.db.Exec("INSERT INTO User (name, surname) VALUES (?,?)", user.Firstname, user.Lastname)
	if err != nil {
		return 0, err
	}
	id, error := query.LastInsertId()
	return int(id), error
}
