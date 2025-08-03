package database

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vickykumar/url_shortner/internal/config"
)

type MysqlDB struct {
	db *sql.DB
}

func NewMySqlConnection(config *config.DatabaseConfig) (*MysqlDB, error) {
	// DSN format: username:password@tcp(host:port)/dbname?parseTime=true

	port, _ := strconv.ParseInt(config.Port, 10, 64)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		port,
		config.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.MaxLifetime)
	db.SetMaxIdleConns(config.MaxIdleConns)

	mysqlDB := &MysqlDB{db: db}
	if err := mysqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	return mysqlDB, nil
}

func (m *MysqlDB) Close() error {
	if m.db != nil {
		if err := m.db.Close(); err != nil {
			return fmt.Errorf("failed to close MySQL connection: %w", err)
		}
	}
	return nil
}

func (m *MysqlDB) Ping() error {
	if m.db == nil {
		return fmt.Errorf("MySQL connection is nil")
	}
	return m.db.Ping()
}

func (m *MysqlDB) GetDB() *sql.DB {
	return m.db
}

func (m *MysqlDB) BeginTx() (*sql.Tx, error) {
	tx, err := m.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin MySQL transaction: %w", err)
	}
	return tx, nil
}
