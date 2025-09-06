package db

import (
	"fmt"

	"github.com/rajanyadav1162/go-microservice-system/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	return DB.AutoMigrate(&model.TicketOrder{})
}
