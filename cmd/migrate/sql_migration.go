package migrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func runMigrationsUp(db *sql.DB) error {

	defer db.Close()
	sqlBytes, err := ioutil.ReadFile("internal/database/migrations/001_create_url.sql")
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	queries := strings.Split(string(sqlBytes), ";")
	for _, q := range queries {
		q = strings.TrimSpace(q)
		if q == "" {
			continue
		}
		if _, err := db.Exec(q); err != nil {
			return fmt.Errorf("failed to execute query: %s, error: %w", q, err)
		}
	}

	log.Println("Migration completed successfully")
	return nil
}
