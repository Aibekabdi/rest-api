package service

import "github.com/Aibekabdi/rest-api/internal/repository"

type Substr interface {
}

type Servise struct {
	Substr
}

func NewService(repos *repository.Repository) *Servise {
	return &Servise{}
}
