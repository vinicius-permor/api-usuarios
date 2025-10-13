package repositories

import (
	"database/sql"
	"log"

	"vinicius-permor/apiGin/src/models"
)

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func (r *UsersRepository) Create(user models.Users) (int64, error) {
	statement := "insert into users (name , email, password) values (?,?,?)"
	result, err := r.db.Exec(statement, user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *UsersRepository) SearchID(id string) (*models.Users, error) {
	statement := "select id , name , email from users where id = ?"

	user := &models.Users{}
	err := r.db.QueryRow(statement, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UsersRepository) UpdateID(id string, user *models.Users) error {
	statement := "update set name = ?, email = ? , where id = ?"
	_, err := r.db.Exec(statement, user.Name, user.Email, id)
	return err
}

func (r *UsersRepository) DeleteUser(id string) error {
	statement := "delete from users where id = ?"
	_, err := r.db.Exec(statement, id)
	return err
}

func (r *UsersRepository) ListallUsers() ([]models.Users, error) {
	statement := "select id, name, email, password from users"

	lineRows, err := r.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := lineRows.Close(); err != nil {
			log.Printf("erro ao fechar o execucao listar todos os usuarios , verifique o erro e tente novamente: %v", err)
		}
	}()
	users := []models.Users{}
	for lineRows.Next() {
		var user models.Users
		err := lineRows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
