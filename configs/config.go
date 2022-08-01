package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Uri      string
	Database string
}

func LoadAPIConfigs() (APIConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg_api, err := getAPI()

	return cfg_api, nil
}

func loadDBConfigs() (DBConfig, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg_db, err := getDB()

	return cfg_db, nil
}

func getAPI() (APIConfig, error) {

	cfg_api := APIConfig{
		Port: os.Getenv("API_PORT"),
	}

	if len(cfg_api.Port) <= 0 {
		log.Fatal("Error api infos in .env file")
	}

	return cfg_api, nil
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
