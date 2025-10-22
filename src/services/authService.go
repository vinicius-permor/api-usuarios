package services

import (
	"errors"
	"vinicius-permor/apiGin/src/models"
	"vinicius-permor/apiGin/src/password"
	"vinicius-permor/apiGin/src/repositories"
)

type AuthService struct {
	userRepo *repositories.UsersRepository
}

func NewAuthService(userRepo *repositories.UsersRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(email, senha string) (*models.Users, error) {
	user, err := s.userRepo.SearchByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario nao econtrado")
	}

	/*
	 * comparacao de senha alterada com hash armazenado
	 */
	if err := password.CheckPasswordHash(senha, user.Password); err != nil {
		return nil, errors.New("senha incorreta")
	}
	user.Password = ""

	return user, nil
}
