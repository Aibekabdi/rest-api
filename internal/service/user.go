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
func (u *UserService) GetUser(id int) (models.UserModel, error) {
	return u.repo.GetUser(id)
}

func (u *UserService) PutUser(id int, user models.UserModel) error {
	return u.repo.PutUser(id, user)
}
func (u *UserService) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
