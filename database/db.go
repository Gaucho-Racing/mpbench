package database

import (
	"fmt"

	singlestore "github.com/singlestore-labs/gorm-singlestore"
	"gorm.io/gorm"
)

func ConnectDB(databaseUser string, databasePassword string, databaseHost string, databasePort string, databaseName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	db, err := gorm.Open(singlestore.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
