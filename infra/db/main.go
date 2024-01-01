package db

import (
	"fmt"
	"go-ddd/util/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DB(logger logger.Logger) (*gorm.DB, error) {
	dsn := "<CONNSTRING>"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Log(fmt.Sprintf("failed initializing database. %v", err))
		return nil, err
	}

	return db, nil
}
