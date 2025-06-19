package storage

import (
	"context"
)

func SaveMessage(senderID, receiverID int, message string) error {
	_, err := DB.Exec(
		context.Background(),
		`INSERT INTO messages (sender_id, receiver_id, message) VALUES ($1, $2, $3)`,
		senderID, receiverID, message,
	)
	return err
}
