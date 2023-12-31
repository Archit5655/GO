package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string 
	Port     string
	Passwrod string
	User     string
	Dbname   string
	SSLModel string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s  dbname=%s  sslmode=%s ", config.Host, config.Port,config.User, config.Passwrod, config.Dbname, config.SSLModel)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err

	}
	return db, nil

}
