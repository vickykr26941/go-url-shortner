package database

import (
	"database/sql"
	"fmt"
	"github.com/vickykumar/url_shortner/internal/config"
	"time"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresConnection(config *config.DatabaseConfig) (*PostgresDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		config.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.MaxLifetime)
	db.SetMaxIdleConns(int(time.Duration(config.MaxIdleConns) * time.Second))

	pgDB := &PostgresDB{db: db}
	if err := pgDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pgDB, nil
}

func (p *PostgresDB) Close() error {
	if p.db != nil {
		if err := p.db.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %w", err)
		}
	}
	return nil
}

func (p *PostgresDB) Ping() error {
	if p.db == nil {
		return fmt.Errorf("database connection is nil")
	}
	return p.db.Ping()
}

func (p *PostgresDB) GetDB() *sql.DB {
	return p.db
}

func (p *PostgresDB) BeginTx() (*sql.Tx, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	return tx, nil
}
