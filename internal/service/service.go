package service

import (
	"github.com/Aibekabdi/rest-api/internal/models"
	"github.com/Aibekabdi/rest-api/internal/repository"
)

type Substr interface {
	LongestSubstrFind(str string) string
}
type Email interface {
	FindEmails(email models.Email) models.Email
}
type Servise struct {
	Substr
	Email
}

func NewService(repos *repository.Repository) *Servise {
	return &Servise{
		Substr: NewSubstrService(),
		Email:  NewEmailServise(),
	}
}
