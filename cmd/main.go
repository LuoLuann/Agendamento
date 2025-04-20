package main

import (
	"log"

	"github.com/LuoLuann/Agendamento/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatalf("Falha ao conectar no DB: %v", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalf("Falha ao obter o objeto sql.DB: %v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Falha no ping ao banco de dados: %v", err)
	}
	log.Println("✅ Banco de dados conectado com sucesso!")

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hehehe funcionou",
		})
	})
	router.Run(":3000")
}
