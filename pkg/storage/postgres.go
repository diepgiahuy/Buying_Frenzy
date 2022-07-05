package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(host string, user string, password string, dbName string, port string) *gorm.DB {
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)
	dsn := fmt.Sprintf("host=%s dbname=%s port=%s", host, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
