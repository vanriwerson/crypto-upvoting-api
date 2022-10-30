package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DbConnectionStr = ""
	Port            = 0
	SecretKey       []byte
)

// LoadEnvVars inicializa as váriaveis de ambiente
func LoadEnvVars() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// DbConnectionStr faz a interpolação de user, password, database_name e as opções necessárias para a conexão com o banco de dados
	DbConnectionStr = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_DB"),
	)

	// Port define a porta onde a api estará acessível
	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 5000
	}

	SecretKey = []byte(os.Getenv("JWT_SECRET"))
}
