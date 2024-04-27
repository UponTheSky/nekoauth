package service

import (
	"fmt"
	"nekoauth/lib/database"
)

// the database is very simple - with only one table and the table has only two columns
func fetchRow(clientId string) (id string, secretHashed string, err error) {
	db, err := database.DB()

	if err != nil {
		return "", "", err
	}

	defer db.Close()

	rows, err := db.Query(`SELECT id, secret_hashed FROM clients WHERE id = $1`, clientId)

	if err != nil {
		return "", "", err
	}

	defer rows.Close()

	if !rows.Next() {
		return "", "", fmt.Errorf("the row with id = %q Not found", clientId)
	}

	if err := rows.Scan(&id, &secretHashed); err != nil {
		return "", "", err
	}

	if rows.Next() {
		return "", "", fmt.Errorf("multiple rows for a single client with id = %q", clientId)
	}

	err = rows.Err()

	return
}
