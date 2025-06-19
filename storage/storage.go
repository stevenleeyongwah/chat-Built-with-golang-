package storage

import (
	"database/sql"
)

var DB *sql.DB

func SaveMessage(senderID, receiverID int, message string) error {
	_, err := DB.Exec(
		`INSERT INTO messages (sender_id, receiver_id, message) VALUES ($1, $2, $3)`,
		senderID, receiverID, message,
	)
	return err
}
