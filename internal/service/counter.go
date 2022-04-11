package service

import (
	"github.com/Aibekabdi/rest-api/internal/repository"
)

type CounterServise struct {
	repo repository.Counter
}

func NewCounterService(repo repository.Counter) *CounterServise {
	return &CounterServise{repo: repo}
}

func (s *CounterServise) Add(num string) error {
	return s.repo.Add(num)
}

func (s *CounterServise) Sub(num string) error {
	return s.repo.Sub(num)
}

func (s *CounterServise) Val() (int, error) {
	return s.repo.Val()
}
