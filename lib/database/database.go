package database

import (
	"database/sql"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func DB() (*sql.DB, error) {
	return sql.Open("sqlite", ":memory:")
}

// bootstrap - this is only for the testing purpose
func BootstrapDB() error {
	db, err := DB()

	if err != nil {
		return err
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS clients (
		id TEXT(64) PRIMARY KEY,
		secret_hashed TEXT(255) NOT NULL
	);`); err != nil {
		return err
	}

	clientId := os.Getenv("CLIENT_ID")
	clientSecretHashed, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("CLIENT_SECRET")), 14)

	if err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO clients (id, secret_hashed) VALUES ($1, $2)`, clientId, clientSecretHashed); err != nil {
		return err
	}

	return nil
}
