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

func (u *UserSqlite) GetUser(id int) (models.UserModel, error) {
	var user = models.UserModel{}
	err := u.db.QueryRow("SELECT name, surname FROM User where id = ?", id).Scan(&user.Firstname, &user.Lastname)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserSqlite) PutUser(id int, user models.UserModel) error {
	if _, err := u.db.Exec("UPDATE User SET name = ?, surname = ? where id = ?", user.Firstname, user.Lastname, id); err != nil {
		return err
	}
	return nil
}

func (u *UserSqlite) DeleteUser(id int) error {
	if _, err := u.db.Exec("DELETE FROM User where id = ?", id); err != nil {
		return err
	}
	return nil
}
