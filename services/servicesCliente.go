package services

import "vinicius-permor/apiGin/repositories"

type UserService struct {
	repo *repositories.UsersRepository
}

func NewUserService(repo *repositories.UsersRepository) *UserService {
	return &UserService{repo: repo}
}
