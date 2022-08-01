package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Uri      string
	Database string
}

func load() (DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg_db, err := getDB()

	return cfg_db, nil
}

func getDB() (DBConfig, error) {

	cfg_db := DBConfig{
		Uri:      os.Getenv("MONGO_URI"),
		Database: os.Getenv("DATABASE_NAME"),
	}

	if len(cfg_db.Uri) <= 0 || len(cfg_db.Database) <= 0 {
		log.Fatal("Error database infos in .env file")
	}

	return cfg_db, nil
}
