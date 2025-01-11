package db

import (
	"log"

	"github.com/killerrekt/Go-Fiber-Auth/internal/model"
)

func RunMigration() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to run the migrations")
	}
	log.Println("Migration completed successfully")
}
