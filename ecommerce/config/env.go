package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DbnName    string
	JWTExpirationInSeconds int64
	JWTSecret	string

}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PublicHost", "http://localhost"),
		Port:       getEnv("Port", "8080"),
		DBUser:     getEnv("DB_User", "joshua468"),
		DBPassword: getEnv("DB_Password", "Temitope2080"),
		DBAddress: fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv(
			"DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "ecom"),
		JWTSecret: getEnv("JWT_SECRET",
		"not-secret-secret-anymore?"),

		JWTExpirationInSeconds:getEnvAsInt("JWT_Exp",3600 *24 *7),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string,fallback int64) int64 {
	if value,ok := os.LookupEnv(key);ok {
		i,err := strconv.ParseInt(value,10,64)
		if err!= nil {
			return fallback
		}
		return i
	}
	return fallback
}
