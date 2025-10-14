package password

import (
	"golang.org/x/crypto/bcrypt"
)

/*
*funcao HashPassword sera exportada para o pacote model sera usada para criar uma senha encriptada para seguraca de acesso,
*somente o usuario que  acessar a senha tem acesso a senha verdadeira
* */
func HashPassword(passKey string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(passKey), bcrypt.DefaultCost)
}

/*
*funcao CheckPasswordHash() sera exportada para o pacote model sera usada para comparar a senha encriptada com a senha inserida,
*pelo  usuario. */
func CheckPasswordHash(passKey, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(passKey))
}
