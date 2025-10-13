package models

// estrutura de Users para receber os dados no formato de JSON, sera exportado para o pacote repositories
type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
