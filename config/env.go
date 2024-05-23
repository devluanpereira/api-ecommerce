package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),                                       // Luan Pereira - aqui define o endereço local
		Port:       getEnv("PORT", "8080"),                                                          // Luan Pereira aqui define a portal local de desenvolvimeno
		DBUser:     getEnv("DB_USER", "root"),                                                       // Luan Pereira - aqui define o usuario no MySQL
		DBPassword: getEnv("DB_PASSWORD", "senha"),                                                  // Luan Pereira - aqui define a senha do usuario mysql
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")), // Luan - Pereira aqui define o endereço do banco e porta
		DBName:     getEnv("DB_NAME", "apiecommerce"),                                               // Luan Pereira - Aqui define o nome Banco de Dados
	}

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
