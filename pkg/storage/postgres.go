package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresStore struct {
	Db *gorm.DB
}

func NewDB(host string, user string, password string, dbName string, port string) *PostgresStore {
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)
	dsn := fmt.Sprintf("host=%s dbname=%s port=%s", host, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &PostgresStore{Db: db}
}

type DbStore interface {
	GetRestaurantStore() *RestaurantStore
	GetMenuStore() *MenuStore
	GetHistoryStore() *HistoryStore
	GetUserStore() *UserStore
}

func (p *PostgresStore) GetRestaurantStore() *RestaurantStore {
	return NewRestaurantStore(p.Db)
}
func (p *PostgresStore) GetMenuStore() *MenuStore {
	return NewMenuStore(p.Db)
}
func (p *PostgresStore) GetHistoryStore() *HistoryStore {
	return NewHistoryStore(p.Db)
}
func (p *PostgresStore) GetUserStore() *UserStore {
	return NewUserStore(p.Db)
}

// WithTx enables repository with transaction
func (p *PostgresStore) WithTx(txHandle *gorm.DB) *PostgresStore {
	if txHandle == nil {
		log.Print("Transaction Database not found")
		return p
	}
	p.Db = txHandle
	return p
}
