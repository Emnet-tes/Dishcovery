package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB ping failed: ", err)
	}

	log.Println("Connected to PostgreSQL DB")
}

type Config struct {
	JWTSecret      string
	CloudinaryURL  string
	ChapaSecretKey string
	HasuraEndpoint string
	HasuraAdminKey string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret:      os.Getenv("JWT_SECRET"),
		CloudinaryURL:  os.Getenv("CLOUDINARY_URL"),
		ChapaSecretKey: os.Getenv("CHAPA_SECRET_KEY"),
		HasuraEndpoint: os.Getenv("HASURA_ENDPOINT"),
		HasuraAdminKey: os.Getenv("HASURA_ADMIN_KEY"),
	}
}
