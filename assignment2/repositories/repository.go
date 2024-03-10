package repositories

import (
	"assignment2/helpers"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	dbConfig := helpers.LoadDatabaseConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Port)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
