package datab

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos (
        id          SERIAL PRIMARY KEY,
        title       VARCHAR(255) NOT NULL,
        completed   BOOLEAN DEFAULT FALSE,
        created_at  TIMESTAMP DEFAULT NOW(),
        finished_at TIMESTAMP NULL
    )`

	_, err := db.Exec(query)
	return err
}
