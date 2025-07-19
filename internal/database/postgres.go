package database

import (
	"database/sql"
	"github.com/vickykumar/url_shortner/internal/config"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresConnection(config *config.DatabaseConfig) (*PostgresDB, error) {
	// TODO: Create PostgreSQL connection
	// TODO: Configure connection pool
	// TODO: Ping database to verify connection
	return nil, nil
}

func (p *PostgresDB) Close() error {
	// TODO: Close database connection
	return nil
}

func (p *PostgresDB) Ping() error {
	// TODO: Ping database
	return nil
}

func (p *PostgresDB) GetDB() *sql.DB {
	// TODO: Return database instance
	return p.db
}

func (p *PostgresDB) BeginTx() (*sql.Tx, error) {
	// TODO: Begin transaction
	return nil, nil
}
