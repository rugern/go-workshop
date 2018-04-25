package database

import (
	"database/sql"
	// Anonymous import - needed for the driver to register itself, and be available to "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/labstack/gommon/log"
)

// Setup creates connection pool, and checks for database connectivity
func Setup(driver string, dataSource string, maxIdleConnections int, maxOpenConnections int) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Errorf("Failed to open database connection: %v", err)
		return nil, err
	}
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(4 * time.Minute) // Must be shorter than wait_time in connection string
	err = db.Ping()
	if err != nil {
		log.Errorf("Failed to ping database connection: %v", err)
		panic(err)
	}

	return db, nil
}
