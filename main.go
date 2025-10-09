package main

import (
	"log"

	"vinicius-permor/apiGin/config"
	"vinicius-permor/apiGin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.Conn()
	if err != nil {
		log.Fatal("erro ao conectar com o banco de dados:", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("erro ao fechar banco de dados , verique o erro e tente novamente: %v", err)
		}
	}()

	r := gin.Default()
	routes.SetupRoutes(r)

	if err := r.Run(":3333"); err != nil {
		log.Printf("erro ao iniciar o sevidor verifique o erro e tente novamente: %v", err)
	}
}
