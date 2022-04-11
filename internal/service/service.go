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
type Servise struct {
	Substr
	Email
	Counter
}

func NewService(repo *repository.Repository) *Servise {
	return &Servise{
		Counter: NewCounterService(repo),
		Substr:  NewSubstrService(),
		Email:   NewEmailServise(),
	}
}
