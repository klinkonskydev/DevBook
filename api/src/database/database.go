package database

import (
	"database/sql"
	"fmt"

	"api/src/config"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect opens connection with database
func Connect() (*sql.DB, error) {
  fmt.Println(config.DataSourceName)
	db, err := sql.Open("mysql", config.DataSourceName)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
    db.Close()
		return nil, err
	}

  return db, nil
}
