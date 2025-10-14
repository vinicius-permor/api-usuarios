package models

import (
	"errors"
	"strings"

	"vinicius-permor/apiGin/src/password"

	"github.com/badoux/checkmail"
)

// estrutura de Users para receber os dados no formato de JSON, sera exportado para o pacote repositories
type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (userModel *Users) Format(step string) error {
	userModel.Name = strings.TrimSpace(userModel.Name)
	userModel.Email = strings.TrimSpace(userModel.Email)
	if step == "update" {
		passHash, err := password.HashPassword(userModel.Password)
		if err != nil {
			return err
		}
		userModel.Password = string(passHash)
	}
	return nil
}

func (userModel *Users) Validade(step string) error {
	if userModel.Name == "" {
		return errors.New("o nome e obrigatorio")
	}

	if userModel.Email == "" {
		return errors.New("o email e obrigatorio")
	}

	if err := checkmail.ValidateFormat(userModel.Email); err != nil {
		return errors.New("email e obrigatorio")
	}

	if step == "update" && userModel.Password == "" {
		return errors.New("a senha e obrigatoria")
	}
	return nil
}

func (userModel *Users) Prepare(step string) error {
	if err := userModel.Validade(step); err != nil {
		return err
	}
	if err := userModel.Validade(step); err != nil {
		return err
	}
	return nil
}
