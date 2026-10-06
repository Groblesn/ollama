package sqlinjection

import (
	"database/sql"
	"fmt"
)

func FindUser(db *sql.DB, username string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT id, email FROM users WHERE username = '%s'", username)
	return db.Query(query)
}
