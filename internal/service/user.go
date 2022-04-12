package service

import (
	"github.com/Aibekabdi/rest-api/internal/models"
	"github.com/Aibekabdi/rest-api/internal/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) PostUserindb(user models.UserModel) (int, error) {
	return u.repo.PostUserindb(user)
}
