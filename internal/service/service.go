package service

import (
	"github.com/Aibekabdi/rest-api/internal/models"
	"github.com/Aibekabdi/rest-api/internal/repository"
)

type Counter interface {
	Add(num string) error
	Sub(num string) error
	Val() (int, error)
}
type Substr interface {
	LongestSubstrFind(str string) string
}
type Email interface {
	FindEmails(email models.Email) models.Email
}
type User interface {
	PostUserindb(user models.UserModel) (int, error)
}
type Servise struct {
	Substr
	Email
	Counter
	User
}

func NewService(repo *repository.Repository) *Servise {
	return &Servise{
		Counter: NewCounterService(repo),
		Substr:  NewSubstrService(),
		Email:   NewEmailServise(),
		User:    NewUserService(repo),
	}
}
