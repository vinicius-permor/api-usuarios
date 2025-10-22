package repositories

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"vinicius-permor/apiGin/src/models"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

//var db UsersRepository

// function Create de criacao de usuarios sera export para o pacote controllers
// a fuction Create esta usando o sqlx com execucao de query no banco de dados sera export no pcate controllers
func (r *UsersRepository) Create(user *models.Users) (int64, error) {

	statement := "insert into users (name , email, password) values (?,?,?)"
	result, err := r.db.Exec(statement, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err

	}
	return id, nil
}

// funcao SearchID de busca de clientes por id pode ser exportada para o pacote controllers
func (r *UsersRepository) SearchID(id string) (*models.Users, error) {
	user := &models.Users{}
	statement := "select id ,name, email from users where id = ?"
	err := r.db.Get(user, statement, id)
	if err == sql.ErrNoRows {
		return nil, errors.New("ususario nao encontrado")
	}
	return user, nil
}

// funcao UpdateID de atualizacao de usuarios sera exportada para o pacote controllers
func (r *UsersRepository) UpdateID(id string, user *models.Users) error {

	statement := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	_, err := r.db.Exec(statement, user.Name, user.Email, user.Password, id)
	return err
}

// funcao DeleteUser que vai deletar o ususario sera exportada para o pacote controllers
func (r *UsersRepository) DeleteUser(id string) error {

	statement := "delete from users where id = ?"
	_, err := r.db.Exec(statement, id)
	return err

}

// func ListallUsers de listagens de todos os ususaios ser exportada para o pacote controllers
func (r *UsersRepository) ListAllUsers() ([]models.Users, error) {
	var user []models.Users
	statement := "select id, name, email, password from users"
	err := r.db.Select(&user, statement)
	return nil, err
}

func (r *UsersRepository) SearchByEmail(email string) (*models.Users, error) {
	user := &models.Users{}
	statement := "SELECT id, name, email, password FROM users WHERE email = ?"
	err := r.db.Get(user, statement, email)
	if err != nil {
		return nil, err

	}
	return user, nil
}
