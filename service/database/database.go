package database

import (
	"contabi-be/config"
	"database/sql"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// DataBaseService contains the database connection
type DataBaseService struct {
	cfg config.Config
	DB  *sql.DB
}

// NewDatabaseService creates a new database service
func NewDatabaseService(cfg config.Config) (*DataBaseService, error) {
	con := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DataBaseService{
		cfg: cfg,
		DB:  db,
	}, nil
}
