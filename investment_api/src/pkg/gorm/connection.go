package gorm

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db            *gorm.DB
	dbConnOnce    sync.Once
	GetConnection = getConnection
)

func UseDefaultConnection() {
	GetConnection = getConnection
}

func getConnection() *gorm.DB {
	dbConnOnce.Do(func() {
		var err error
		cfg := &gorm.Config{
			PrepareStmt: true,
		}
		db, err = gorm.Open(sqlite.Open("gorm.db"), cfg)

		if err != nil {
			panic("base_plain_handler: error: gorm db not found")
		}

		crearTabla := `
		CREATE TABLE IF NOT EXISTS statistics (
			success BOOLEAN,
			investment INTEGER
		);
		`
		tx := db.Exec(crearTabla)
		if tx.Error != nil {
			log.Fatal(tx.Error)
		}
	})

	return db
}

func Close() error {
	return nil
}
