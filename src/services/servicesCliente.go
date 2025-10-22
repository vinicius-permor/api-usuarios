package services

import (
	"errors"

	"vinicius-permor/apiGin/src/models"
	"vinicius-permor/apiGin/src/repositories"
)

type UserService struct {
	repo *repositories.UsersRepository
}

func NewUserService(repo *repositories.UsersRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.Users) (int64, error) {
	if user.Name == "" {
		return 0, errors.New("usuario nao pode ser nulo")
	}
	if user.Email == "" {
		return 0, errors.New("Email e obrigatorio")
	}
	if len(user.Name) < 3 {
		return 0, errors.New("nome tem que ter ao menos 3 caracteres")
	}
	return s.repo.Create(user)

}

func (s *UserService) SearchUserID(id string) (*models.Users, error) {
	user, err := s.repo.SearchID(id)
	if err != nil {
		return nil, err

	}
	if user == nil {
		return nil, errors.New("usuario nao encontrado")
	}
	return user, nil
}

func (s *UserService) UpdateUserID(id string, user *models.Users) error {
	user, err := s.repo.SearchID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("usuario nao encontrado")
	}
	return s.repo.UpdateID(id, user)
}

func (s *UserService) DeleteUserID(id string) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) ListAllUsers() ([]models.Users, error) {
	user, err := s.repo.ListAllUsers()
	if err != nil {
		return nil, err

	}
	if len(user) == 0 {
		return []models.Users{}, nil
	}
	return user, nil
}
