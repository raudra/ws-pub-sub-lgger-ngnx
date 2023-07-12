package config

import (
	"profile-service/src/models"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	var err error

	log.Print("Setting up DB")

	dbUrl := "postgres://postgres:postgres@host.docker.internal:5435/profile_db"

	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal().
			Err(err).
			Msgf("Error while connecting to DB")

	}

	migrate()
}

func DbClient() *gorm.DB {
	return db
}

func migrate() {
	log.Print("DB migration started")
	db.AutoMigrate(&models.User{})
	log.Print("DB migration end")
}
