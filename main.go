package main

import (
	"log"

	"vinicius-permor/apiGin/src/config"
	"vinicius-permor/apiGin/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.Conn()
	if err != nil {
		log.Fatal("erro ao conectar com banco de dados", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("erro ao fechar banco de dados: %v", err)
		}
	}()
	r := gin.Default()
	if err := r.SetTrustedProxies([]string{"192.168.1.2"}); err != nil {
		log.Printf("erro na checagem de proxy: %v", err)
	}
	routes.SetupRoutes(r, db)

	if err := r.Run(":3333"); err != nil {
		log.Printf("erro ao iniciar o sevidor verifique o erro e tente novamente: %v", err)
	}
}
