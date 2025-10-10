package services

import (
	"errors"

	"vinicius-permor/apiGin/models"
	"vinicius-permor/apiGin/repositories"
)

type UserService struct {
	repo *repositories.UsersRepository
}

func NewUserService(repo *repositories.UsersRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.Users) (int64, error) {
	return s.repo.Create(*user)
}

func (s *UserService) SearchUserID(id string) (*models.Users, error) {
	user, err := s.repo.SearchID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("usuario nao econtrado")
	}
	return user, nil
}

func (s *UserService) UpdateUserID(id string, user *models.Users) error {
	users, err := s.repo.SearchID(id)
	if err != nil {
		return err
	}
	if users == nil {
		return errors.New("usuario nao econtrado")
	}
	return s.repo.UpdateID(id, user)
}

func (s *UserService) DeleteUserID(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) ListAllUsers() ([]models.Users, error) {
	return s.repo.ListallUsers()
}
